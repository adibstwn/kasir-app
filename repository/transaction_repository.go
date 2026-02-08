package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kasir-app/model"
	"time"
)

type ITransactionRepository interface {
	Insert(tx *gorm.DB, transaction model.Transaction) error
	CountTransactionByDate(startTime, endTime time.Time) (int64, error)
}
type TransactionRepositoryImpl struct {
	gormDb *gorm.DB
}

func NewTransactionRepository(gormDb *gorm.DB) ITransactionRepository {
	return &TransactionRepositoryImpl{gormDb: gormDb}
}

func (t TransactionRepositoryImpl) Insert(tx *gorm.DB, transaction model.Transaction) error {
	return tx.Create(&transaction).Error
}

func (t TransactionRepositoryImpl) CountTransactionByDate(startTime, endTime time.Time) (int64, error) {
	var response int64

	err := t.gormDb.Model(&model.Transaction{}).Where("created_at BETWEEN ? AND ?", startTime, endTime).Count(&response).Error
	if err != nil {
		fmt.Println("error count total transaction today ", err)
		return 0, err
	}
	return response, nil
}
