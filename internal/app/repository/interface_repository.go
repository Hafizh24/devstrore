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

type IUserRepository interface {
	Create(user model.User) error
	GetByEmail(email string) (model.User, error)
	GetByID(id int) (model.User, error)
}
type IAuthRepository interface {
	Create(auth model.Auth) error
	Find(userID int, RefreshToken string) (model.Auth, error)
	Delete(userID int) (model.Auth, error)
}
