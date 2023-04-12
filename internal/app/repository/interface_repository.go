package repository

import "github.com/hafizh24/devstore/internal/app/model"

type ICategoryRepository interface {
	Browse() ([]model.Category, error)
	Create(category model.Category) error
	GetByID(id string) (model.Category, error)
	Update(id string, category model.Category) error
	Delete(id string) (model.Category, error)
}

type IProductRepository interface {
	Browse() ([]model.Product, error)
	GetByID(id string) (model.Product, error)
	Create(product model.Product) error
	Delete(id string) (model.Product, error)
	Update(id string, product model.Product) error
}
