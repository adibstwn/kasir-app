package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"kasir-app/model"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll() ([]model.Product, error) {
	query := "SELECT p.id, p.name, p.price, p.stock, p.is_deleted, c.name as category_name FROM product p left join category c on p.category_id = c.id where p.is_deleted = 'N'"
	rows, err := repo.db.Query(query)
	if err != nil {
		fmt.Println("error get all product ", err)
		return nil, err
	}
	defer rows.Close()
	products := make([]model.Product, 0)
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.IsDeleted, &product.CategoryName)
		if err != nil {
			log.Println("error scan data ", err)

			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *ProductRepository) Create(product *model.Product) error {
	query := "INSERT INTO product (id, name, price, stock, category_id, is_deleted) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := repo.db.QueryRow(query, product.Id, product.Name, product.Price, product.Stock, product.IdCategory, product.IsDeleted).Scan(&product.Id)
	return err
}

func (repo *ProductRepository) GetByID(id string) (*model.Product, error) {
	query := "SELECT id, name, price, stock, is_deleted FROM product WHERE id = $1 and is_deleted = 'N'"

	var p model.Product
	err := repo.db.QueryRow(query, id).Scan(&p.Id, &p.Name, &p.Price, &p.Stock, &p.IsDeleted)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}
func (repo *ProductRepository) Update(product *model.Product) error {
	query := "UPDATE product SET name = $1, price = $2, stock = $3 WHERE id = $4"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.Id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (repo *ProductRepository) Delete(id string) error {
	query := "UPDATE product SET is_deleted='Y' WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		//errors.Is(err, sql.ErrNoRows)
		return err
	}

	return err
}
