package repository

import "github.com/hafizh24/devstore/internal/app/model"

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
	Delete(userID int) error
}
