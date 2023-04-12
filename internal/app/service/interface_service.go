package service

import (
	"github.com/hafizh24/devstore/internal/app/schema"
)

type ICategoryService interface {
	BrowseAll() ([]schema.GetCategoryResp, error)
	Create(req *schema.CreateCategoryReq) error
	GetByID(id string) (schema.GetCategoryResp, error)
	UpdateByID(id string, req *schema.UpdateCategoryReq) error
	DeleteByID(id string) (schema.GetCategoryResp, error)
}

type IProductService interface {
	BrowseAll() ([]schema.GetProductResp, error)
	GetByID(id string) (schema.GetProductResp, error)
	Create(req *schema.CreateProductReq) error
	DeleteByID(id string) (schema.GetProductResp, error)
	UpdateByID(id string, req *schema.UpdateProductReq) error
}
