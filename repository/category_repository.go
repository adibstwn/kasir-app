package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"kasir-app/model"
	"log"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAll() ([]model.Category, error) {
	rows, err := repo.db.Query("select id, name, description, is_delete from category where is_delete='N'")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var categories = make([]model.Category, 0)
	for rows.Next() {
		var data model.Category
		err = rows.Scan(&data.Id, &data.Name, &data.Description, &data.IsDelete)
		if err != nil {
			log.Println("error scan data ", err)
			return nil, err
		}
		categories = append(categories, data)
	}
	return categories, nil
}
func (repo *CategoryRepository) GetById(id string) (*model.Category, error) {
	query := "SELECT id, name, description, is_delete from category where id=$1"

	var c model.Category
	err := repo.db.QueryRow(query, id).Scan(&c.Id, &c.Name, &c.Description, &c.IsDelete)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return &c, nil
}

func (repo *CategoryRepository) Create(category *model.Category) error {
	query := "INSERT INTO category(id,name,description,is_delete) VALUES ($1,$2,$3,$4)"
	_, err := repo.db.Exec(query, category.Id, category.Name, category.Description, category.IsDelete)
	if err != nil {
		return err
	}
	return nil
}
func (repo *CategoryRepository) Update(category *model.Category) error {
	query := "UPDATE category SET name = $1, description = $2 WHERE id = $3"
	_, err := repo.db.Exec(query, category.Name, category.Description, category.Id)
	if err != nil {
		log.Println("error update category ", err)
		return err
	}
	return nil
}
func (repo *CategoryRepository) Delete(id string) error {
	query := "UPDATE category set is_delete='Y' WHERE id = $1"
	_, err := repo.db.Exec(query, id)
	if err != nil {
		log.Println("error delete category ", err)
		return err
	}
	return nil
}
