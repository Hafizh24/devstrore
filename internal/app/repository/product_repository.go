package repository

import (
	"errors"
	"fmt"

	"github.com/hafizh24/devstore/internal/app/model"
	"github.com/hafizh24/devstore/internal/pkg/reason"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (pr *ProductRepository) Browse() ([]model.Product, error) {

	var products []model.Product
	var sqlStatement = `
		SELECT id,name,description, currency, price, total_stock, is_active, category_id
		FROM products
	`
	rows, err := pr.DB.Queryx(sqlStatement)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Browse : %w", err))
		return products, err
	}

	for rows.Next() {
		var product model.Product
		// nolint:errcheck
		rows.StructScan(&product)
		products = append(products, product)
	}
	return products, nil
}

func (pr *ProductRepository) GetByID(id string) (model.Product, error) {
	var product model.Product
	var sqlStatement = `
	SELECT id,name,description, currency, price, total_stock, is_active, category_id
	FROM products
	WHERE id = $1
	`
	err := pr.DB.QueryRowx(sqlStatement, id).StructScan(&product)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - GetByID : %w", err))
		return product, err
	}

	return product, nil
}

func (pr *ProductRepository) Create(product model.Product) error {
	var sqlStatement = `
		INSERT INTO products (name, description, currency, price, total_stock, is_active, category_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		`

	_, err := pr.DB.Exec(sqlStatement, product.Name, product.Description, product.Currency, product.Price, product.TotalStock, product.IsActive, product.CategoryID)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Create : %w", err))
		return err
	}
	return nil
}

func (pr *ProductRepository) Delete(id string) (model.Product, error) {
	var product model.Product
	var sqlStatement = `
	DELETE FROM products
	WHERE id = $1
	`

	_, err := pr.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Delete : %w", err))
		return product, err
	}

	// check, _ := delete.RowsAffected()
	// if check == 0 {
	// 	return product, errors.New(reason.CategoryNotFound)
	// }

	return product, nil
}

func (pr *ProductRepository) Update(id string, product model.Product) error {

	var sqlStatement = `
	UPDATE products 
	SET name= $2, description= $3, currency= $4, price= $5, total_stock= $6, is_active= $7, category_id= $8
	WHERE id = $1
	`
	update, err := pr.DB.Exec(sqlStatement, id, product.Name, product.Description, product.Currency, product.Price, product.TotalStock, product.IsActive, product.CategoryID)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Update : %w", err))
		return err
	}

	check, _ := update.RowsAffected()
	if check == 0 {
		return errors.New(reason.ProductNotFound)
	}

	return nil
}
