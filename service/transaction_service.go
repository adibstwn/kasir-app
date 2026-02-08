package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kasir-app/constant"
	"kasir-app/dto"
	"kasir-app/errors"
	"kasir-app/model"
	"kasir-app/repository"
	"time"
)

type ITransactionService interface {
	Checkout(context context.Context, request *dto.CheckoutRequest) error
	ReportToday() (*dto.TransactionReport, error)
	ReportByDate(startDate, endDate string) (*dto.TransactionReport, error)
}

type TransactionServiceImpl struct {
	transactionRepo       repository.ITransactionRepository
	transactionDetailRepo repository.ITransactionDetailRepository
	productRepo           repository.IProductRepository
	gormDb                *gorm.DB
}

func (t TransactionServiceImpl) ReportByDate(startDate, endDate string) (*dto.TransactionReport, error) {
	var response *dto.TransactionReport

	startDate += " 00:00:00"
	endDate += " 23:59:59"

	setupDate, err := time.Parse(constant.FormatDateTime, startDate)
	if err != nil {
		fmt.Println("error parsing date")
		return nil, err
	}

	// 2. Truncate the time to the start of the day (midnight)
	// This sets the hour, minute, second, and nanosecond to zero
	startDay := time.Date(setupDate.Year(), setupDate.Month(), setupDate.Day(), 0, 0, 0, 0, setupDate.Location())

	// 3. Calculate the end of the day by adding one day to the start of the day
	endDay, err := time.Parse(constant.FormatDateTime, endDate)

	fmt.Printf("Start of today: %v\n", startDay)
	fmt.Printf("End of today (start of tomorrow): %v\n", endDay)
	totalTransaction, err := t.transactionRepo.CountTransactionByDate(startDay, endDay)
	data, err := t.transactionDetailRepo.ReportToday(startDay, endDay)
	if err != nil {
		fmt.Println("error report today ", err)
		return response, err
	}
	records := data

	var totalAmountToday float64
	highestProductName := ""
	tempProductSold := make(map[string]float64)

	for _, record := range records {
		if _, exist := tempProductSold[record.ProductName]; !exist {
			tempProductSold[record.ProductName] = record.Quantity
		} else {
			tempProductSold[record.ProductName] += record.Quantity
		}

		totalAmountToday += record.Amount
	}

	var start float64
	for productName, f := range tempProductSold {
		if f > start {
			start = f
			highestProductName = productName
		}
	}

	response = &dto.TransactionReport{
		TotalRevenue:     totalAmountToday,
		TotalTransaction: totalTransaction,
		BestSellerProduct: dto.BestSellerProduct{
			Name:    highestProductName,
			QtySell: tempProductSold[highestProductName],
		},
	}

	return response, nil
}

func NewTransactionServiceImpl(gormDb *gorm.DB, transactionRepo repository.ITransactionRepository, transactionDetailRepo repository.ITransactionDetailRepository, productRepo repository.IProductRepository) ITransactionService {
	return &TransactionServiceImpl{
		gormDb:                gormDb,
		transactionRepo:       transactionRepo,
		transactionDetailRepo: transactionDetailRepo,
		productRepo:           productRepo,
	}
}

func (t TransactionServiceImpl) Checkout(context context.Context, request *dto.CheckoutRequest) error {

	transactionId := uuid.New().String()
	var totalAmount float64
	return t.gormDb.WithContext(context).Transaction(func(tx *gorm.DB) error {
		var transactionDetails []*model.TransactionDetail
		//var updateProductsData []*model.Product
		for _, item := range request.Items {
			// decrease product
			product, err := t.productRepo.GetByID(tx, item.ProductId)
			if err != nil {
				return err
			}
			//updateProductsData = append(updateProductsData, product)

			productData := &model.Product{
				Id:    product.Id,
				Stock: product.Stock - item.Quantity,
			}
			if productData.Stock <= 0 {
				return errors.NotFound("product stock is not enough")
			}
			fmt.Println("Update product")

			err = t.productRepo.Update(tx, productData) // this will be make N query
			if err != nil {
				return err
			}
			fmt.Println("End update product")

			// create transaction_detail
			transactionDetail := model.TransactionDetail{
				Id:            uuid.New().String(),
				TransactionId: transactionId,
				ProductId:     item.ProductId,
				Quantity:      item.Quantity,
				Amount:        product.Price * item.Quantity,
			}
			totalAmount += transactionDetail.Amount
			transactionDetails = append(transactionDetails, &transactionDetail)
		}
		//batch insert transaction detail
		fmt.Println("Start batch insert transaction detail")
		err := t.transactionDetailRepo.BatchInsert(tx, transactionDetails)
		if err != nil {
			return err
		}

		// insert into transaction
		transaction := model.Transaction{
			Id:          transactionId,
			TotalAmount: totalAmount,
			CreatedAt:   time.Now(),
		}
		err = t.transactionRepo.Insert(tx, transaction)
		if err != nil {
			fmt.Println("error insert transaction", err)
			return err
		}
		return nil
	})
}

func (t TransactionServiceImpl) ReportToday() (*dto.TransactionReport, error) {
	// 1. Get the current time
	var response *dto.TransactionReport
	setupDate := time.Now()

	// 2. Truncate the time to the start of the day (midnight)
	// This sets the hour, minute, second, and nanosecond to zero
	startDay := time.Date(setupDate.Year(), setupDate.Month(), setupDate.Day(), 0, 0, 0, 0, setupDate.Location())

	// 3. Calculate the end of the day by adding one day to the start of the day
	endDay := startDay.AddDate(0, 0, 1)

	fmt.Printf("Start of today: %v\n", startDay)
	fmt.Printf("End of today (start of tomorrow): %v\n", endDay)
	totalTransaction, err := t.transactionRepo.CountTransactionByDate(startDay, endDay)
	data, err := t.transactionDetailRepo.ReportToday(startDay, endDay)
	if err != nil {
		fmt.Println("error report today ", err)
		return response, err
	}
	records := data

	var totalAmountToday float64
	highestProductName := ""
	tempProductSold := make(map[string]float64)

	for _, record := range records {
		if _, exist := tempProductSold[record.ProductName]; !exist {
			tempProductSold[record.ProductName] = record.Quantity
		} else {
			tempProductSold[record.ProductName] += record.Quantity
		}

		totalAmountToday += record.Amount
	}

	var start float64
	for productName, f := range tempProductSold {
		if f > start {
			start = f
			highestProductName = productName
		}
	}

	response = &dto.TransactionReport{
		TotalRevenue:     totalAmountToday,
		TotalTransaction: totalTransaction,
		BestSellerProduct: dto.BestSellerProduct{
			Name:    highestProductName,
			QtySell: tempProductSold[highestProductName],
		},
	}

	return response, nil
}
