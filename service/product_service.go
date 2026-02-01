package service

import (
	"fmt"
	"kasir-app/dto"
	"kasir-app/model"
	"kasir-app/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (product *ProductService) GetAllProduct() []model.Product {
	var products []model.Product
	products, err := product.repo.GetAll()
	if err != nil {
		fmt.Println("error fetching products", err.Error())
	}
	return products

}

func (product *ProductService) GetById(id string) (*model.Product, error) {
	response, err := product.repo.GetByID(id)
	if err != nil {
		fmt.Println("error fetching product", err.Error())
		return nil, err
	}
	return response, nil
}

func (product *ProductService) Create(data *dto.CreateProduct) error {
	productModel := &model.Product{
		Id:         uuid.New().String(),
		Name:       data.Name,
		Price:      data.Price,
		Stock:      data.Stock,
		IdCategory: data.CategoryId,
		IsDeleted:  "N",
	}

	err := product.repo.Create(productModel)
	if err != nil {
		fmt.Println("error creating product", err.Error())
		return err
	}
	return nil
}

func (product *ProductService) Update(data *dto.UpdateProduct) error {
	_, err := product.repo.GetByID(data.ID)
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

	err = product.repo.Update(productModel)
	if err != nil {
		fmt.Println("error updating product", err.Error())
		return err
	}
	return nil
}

func (product *ProductService) Delete(id string) error {
	err := product.repo.Delete(id)
	if err != nil {
		fmt.Println("error deleting product", err.Error())
		return err
	}
	return nil
}
