package repository

import (
	"gorm.io/gorm"
	"kasir-app/model"
	"time"
)

type ITransactionDetailRepository interface {
	InsertTransactionDetail(transactionDetail *model.TransactionDetail) error
	BatchInsert(tx *gorm.DB, transactionDetails []*model.TransactionDetail) error
	ReportToday(startDate, endDate time.Time) ([]model.TransactionReportDb, error)
}
type TransactionDetailRepositoryImpl struct {
	gormDb *gorm.DB
}

func NewTransactionDetailRepository(gormDb *gorm.DB) ITransactionDetailRepository {
	return &TransactionDetailRepositoryImpl{gormDb: gormDb}
}

func (t TransactionDetailRepositoryImpl) InsertTransactionDetail(transactionDetail *model.TransactionDetail) error {

	err := t.gormDb.Create(&transactionDetail).Error
	return err
}

func (t TransactionDetailRepositoryImpl) BatchInsert(tx *gorm.DB, transactionDetails []*model.TransactionDetail) error {

	err := tx.CreateInBatches(&transactionDetails, 500).Error

	return err
}

func (t TransactionDetailRepositoryImpl) ReportToday(startDate, endDate time.Time) ([]model.TransactionReportDb, error) {
	var report []model.TransactionReportDb
	err := t.gormDb.Table("transaction_details td").
		Select("td.*, p.name as product_name ").Joins("left join transactions t on td.transaction_id = t.id").
		Joins("left join products p on td.product_id = p.id").
		Where(" t.created_at between ? and ?", startDate, endDate).Scan(&report).Error
	return report, err
}
