package service

import (
	"kasir-app/dto"
	"kasir-app/model"
	"kasir-app/repository"
	"log"

	"github.com/google/uuid"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}
func (service *CategoryService) Create(category *dto.CreateCategoryReq) error {
	categoryModel := model.Category{
		Id:          uuid.New().String(),
		Name:        category.Name,
		Description: category.Description,
		IsDelete:    "N",
	}
	err := service.repo.Create(&categoryModel)
	if err != nil {
		log.Println("error create category ", err)
	}
	return nil
}

func (service *CategoryService) Update(category *dto.UpdateCategoryReq) error {
	_, err := service.repo.GetById(category.Id)
	if err != nil {
		return err
	}
	categoryModel := model.Category{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
	}
	err = service.repo.Update(&categoryModel)
	if err != nil {
		log.Println("error update category ", err)
		return err
	}
	return nil
}

func (service *CategoryService) GetAll() ([]model.Category, error) {
	var categories []model.Category
	categories, err := service.repo.GetAll()
	if err != nil {
		log.Println("error get all category ", err)
		return nil, err
	}
	return categories, nil
}

func (service *CategoryService) GetById(id string) (*model.Category, error) {
	category, err := service.repo.GetById(id)
	if err != nil {
		log.Println("error get category ", err)
		return nil, err
	}
	return category, nil
}
func (service *CategoryService) DeleteById(id string) error {
	_, err := service.repo.GetById(id)
	if err != nil {
		return err
	}
	err = service.repo.Delete(id)
	if err != nil {
		log.Println("error delete category ", err)
		return err
	}
	return nil
}
