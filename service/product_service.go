package service

import (
	"fmt"
	"github.com/google/uuid"
	"kasir-app/constant"
	"kasir-app/database"
	"kasir-app/dto"
	"kasir-app/model"
	"kasir-app/repository"
)

type IProductService interface {
	GetAllProduct(productName string, page int, maxRow int) (dto.BaseResponse, error)
	GetById(id string) (*model.Product, error)
	Create(data *dto.CreateProduct) error
	Update(data *dto.UpdateProduct) error
	Delete(id string) error
}

type ProductServiceImpl struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) IProductService {
	return &ProductServiceImpl{repo: repo}
}

func (product *ProductServiceImpl) GetAllProduct(productName string, page int, maxRow int) (dto.BaseResponse, error) {
	var response dto.BaseResponse

	products, totalData, err := product.repo.GetAll(productName, page, maxRow)
	if err != nil {
		fmt.Println("errors fetching products", err.Error())
	}

	// calculate total page
	totalPage := (totalData + maxRow - 1) / maxRow

	fmt.Println(" totalPage ", totalPage)
	paging := dto.Pagination{
		Page:      page,
		MaxRow:    maxRow,
		TotalPage: totalPage,
		TotalData: totalData,
	}

	response.StatusCode = constant.SUCCESS
	response.Data = products
	response.Pagination = &paging
	return response, err
}

func (product *ProductServiceImpl) GetById(id string) (*model.Product, error) {
	response, err := product.repo.GetByID(database.GormDB, id)
	if err != nil {
		fmt.Println("errors fetching product", err.Error())
		return nil, err
	}
	return response, nil
}

func (product *ProductServiceImpl) Create(data *dto.CreateProduct) error {
	productModel := &model.Product{
		Id:         uuid.New().String(),
		Name:       data.Name,
		Price:      data.Price,
		Stock:      data.Stock,
		IdCategory: &data.CategoryId,
		IsDeleted:  "N",
	}

	err := product.repo.Create(productModel)
	if err != nil {
		fmt.Println("errors creating product", err.Error())
		return err
	}
	return nil
}

func (product *ProductServiceImpl) Update(data *dto.UpdateProduct) error {
	_, err := product.repo.GetByID(database.GormDB, data.ID)
	if err != nil {
		return err
	}

	productModel := &model.Product{
		Id:        data.ID,
		Name:      data.Name,
		Price:     data.Price,
		Stock:     data.Stock,
		IsDeleted: "N",
	}

	err = product.repo.Update(database.GormDB, productModel)
	if err != nil {
		fmt.Println("errors updating product", err.Error())
		return err
	}
	return nil
}

func (product *ProductServiceImpl) Delete(id string) error {
	//err := product.repo.Delete(id)
	//if err != nil {
	//	fmt.Println("errors deleting product", err.Error())
	//	return err
	//}
	return nil
}
