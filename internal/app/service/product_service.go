package service

import (
	"errors"
	"strconv"

	"github.com/hafizh24/devstore/internal/app/model"
	"github.com/hafizh24/devstore/internal/app/repository"
	"github.com/hafizh24/devstore/internal/app/schema"
	"github.com/hafizh24/devstore/internal/pkg/reason"
)

type ProductService struct {
	productRepo  repository.IProductRepository
	categoryRepo repository.ICategoryRepository
}

func NewProductService(productRepo repository.IProductRepository, categoryRepo repository.ICategoryRepository) *ProductService {
	return &ProductService{productRepo: productRepo, categoryRepo: categoryRepo}
}

func (ps *ProductService) BrowseAll() ([]schema.GetProductResp, error) {
	var resp []schema.GetProductResp

	products, err := ps.productRepo.Browse()
	if err != nil {
		return nil, errors.New(reason.ProductCannotBrowse)
	}

	for _, value := range products {
		var respData schema.GetProductResp
		respData.ID = value.ID
		respData.Name = value.Name
		respData.Description = value.Description
		respData.Currency = value.Currency
		respData.Price = value.Price
		respData.TotalStock = value.TotalStock
		respData.IsActive = value.IsActive
		respData.CategoryID = value.CategoryID
		resp = append(resp, respData)
	}

	return resp, nil
}

func (ps *ProductService) GetByID(id string) (schema.GetDetailResp, error) {
	var resp schema.GetDetailResp

	product, err := ps.productRepo.GetByID(id)
	if err != nil {
		return resp, errors.New(reason.ProductCannotGetDetail)
	}
	categoryID := strconv.Itoa(product.CategoryID)
	category, err := ps.categoryRepo.GetByID(categoryID)
	if err != nil {
		return resp, errors.New(reason.CategoryCannotGetDetail)
	}

	resp.ID = product.ID
	resp.Name = product.Name
	resp.Description = product.Description
	resp.Currency = product.Currency
	resp.Price = product.Price
	resp.TotalStock = product.TotalStock
	resp.IsActive = product.IsActive
	resp.Category = schema.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return resp, nil
}

func (ps *ProductService) Create(req *schema.CreateProductReq) error {
	var insertData model.Product

	insertData.Name = req.Name
	insertData.Description = req.Description
	insertData.Currency = req.Currency
	insertData.Price = req.Price
	insertData.TotalStock = req.TotalStock
	insertData.IsActive = req.IsActive
	insertData.CategoryID = req.CategoryID

	err := ps.productRepo.Create(insertData)
	if err != nil {
		return errors.New(reason.ProductCannotCreate)
	}
	return nil
}

func (ps *ProductService) DeleteByID(id string) (schema.GetProductResp, error) {
	var req schema.GetProductResp

	check, _ := ps.productRepo.GetByID(id)
	if check.ID == 0 {
		return req, errors.New(reason.ProductNotFound)
	}

	product, err := ps.productRepo.Delete(id)
	if err != nil {
		return req, errors.New(reason.ProductCannotDelete)
	}

	req.ID = product.ID
	req.Name = product.Name
	req.Description = product.Description
	req.Currency = product.Currency
	req.Price = product.Price
	req.TotalStock = product.TotalStock
	req.IsActive = product.IsActive
	req.CategoryID = product.CategoryID

	return req, nil
}

func (ps *ProductService) UpdateByID(id string, req *schema.UpdateProductReq) error {
	var updateData model.Product

	updateData.Name = req.Name
	updateData.Description = req.Description
	updateData.Currency = req.Currency
	updateData.Price = req.Price
	updateData.TotalStock = req.TotalStock
	updateData.IsActive = req.IsActive
	updateData.CategoryID = req.CategoryID

	check, _ := ps.productRepo.GetByID(id)
	if check.ID == 0 {
		return errors.New(reason.ProductNotFound)
	}

	err := ps.productRepo.Update(id, updateData)
	if err != nil {
		return errors.New(reason.ProductCannotUpdate)
	}

	return nil
}
