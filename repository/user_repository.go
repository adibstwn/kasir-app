package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kasir-app/model"
)

type UserRepository struct {
	gormDb *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) *UserRepository {
	return &UserRepository{gormDb}
}

func (repo *UserRepository) GetAll() (*[]model.User, error) {
	return nil, nil
}
func (repo *UserRepository) GetById(id uint) (*model.User, error) {
	return nil, nil
}

func (repo *UserRepository) Create(user *model.User) error {

	if err := repo.gormDb.Create(&user).Error; err != nil {
		fmt.Println("errors create user")
		return err
	}
	return nil
}
