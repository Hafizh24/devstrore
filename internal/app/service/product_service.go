package service

import (
	"errors"

	"github.com/hafizh24/devstore/internal/app/model"
	"github.com/hafizh24/devstore/internal/app/repository"
	"github.com/hafizh24/devstore/internal/app/schema"
	"github.com/hafizh24/devstore/internal/pkg/reason"
)

type ProductService struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (ps *ProductService) BrowseAll() ([]schema.GetProductResp, error) {
	var resp []schema.GetProductResp

	products, err := ps.repo.Browse()
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

func (ps *ProductService) GetByID(id string) (schema.GetProductResp, error) {
	var resp schema.GetProductResp

	product, err := ps.repo.GetByID(id)
	if err != nil {
		return resp, errors.New(reason.ProductCannotGetDetail)
	}

	resp.ID = product.ID
	resp.Name = product.Name
	resp.Description = product.Description
	resp.Currency = product.Currency
	resp.Price = product.Price
	resp.TotalStock = product.TotalStock
	resp.IsActive = product.IsActive
	resp.CategoryID = product.CategoryID

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

	err := ps.repo.Create(insertData)
	if err != nil {
		return errors.New(reason.ProductCannotCreate)
	}
	return nil
}

func (ps *ProductService) DeleteByID(id string) (schema.GetProductResp, error) {
	var req schema.GetProductResp

	product, err := ps.repo.Delete(id)
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

	if req.ID == 0 {
		return req, errors.New(reason.ProductNotFound)
	}

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

	if updateData.ID == 0 {
		return errors.New(reason.ProductNotFound)
	}

	err := ps.repo.Update(id, updateData)
	if err != nil {
		return errors.New(reason.ProductCannotUpdate)
	}

	return nil
}
