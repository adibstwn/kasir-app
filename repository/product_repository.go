package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kasir-app/model"
)

type IProductRepository interface {
	GetAll(productName string, page int, maxRow int) ([]model.Product, int, error)
	GetByID(tx *gorm.DB, id string) (*model.Product, error)
	Create(product *model.Product) error
	Update(tx *gorm.DB, product *model.Product) error
	BatchUpdate(tx *gorm.DB, products []*model.Product) error
}
type ProductRepository struct {
	gormDb *gorm.DB
}

func NewProductRepository(gormDb *gorm.DB) IProductRepository {
	return &ProductRepository{
		gormDb: gormDb,
	}
}

func (repo *ProductRepository) GetAll(productName string, page int, maxRow int) ([]model.Product, int, error) {
	// setup pagination
	if page < 1 {
		page = 1
	}
	if maxRow > 100 {
		maxRow = 100
	}

	var products []model.Product
	var total int64

	db := repo.gormDb.Model(&model.Product{})

	// start filter
	if productName != "" {
		searchQuery := fmt.Sprintf("%%%s%%", productName)
		db = db.Where("name LIKE ?", searchQuery)
	}
	db = db.Where("is_deleted = 'N'")
	// end filter

	// count first
	if err := db.Count(&total).Error; err != nil {
		fmt.Println("error count products")
		return nil, 0, err
	}
	// get data
	offset := (page - 1) * maxRow
	err := db.
		Limit(maxRow).
		Offset(offset).
		Find(&products).
		Error

	return products, int(total), err
}

func (repo *ProductRepository) Create(product *model.Product) error {
	err := repo.gormDb.Create(&product).Error

	return err
}

func (repo *ProductRepository) GetByID(tx *gorm.DB, id string) (*model.Product, error) {
	var product *model.Product
	err := tx.Find(&product, "id = ?", id).Error

	return product, err

}

func (repo *ProductRepository) Update(tx *gorm.DB, product *model.Product) error {
	return tx.Updates(&product).Error
}

func (repo *ProductRepository) Delete(id string) error {
	err := repo.gormDb.Model(&model.Product{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": "Y",
		}).Error

	return err
}

func (repo *ProductRepository) BatchUpdate(tx *gorm.DB, products []*model.Product) error {
	err := tx.Save(&products).Error
	return err
}
