package service

import (
	"errors"

	"github.com/hafizh24/devstore/internal/app/model"
	"github.com/hafizh24/devstore/internal/app/repository"
	"github.com/hafizh24/devstore/internal/app/schema"
	"github.com/hafizh24/devstore/internal/pkg/reason"
)

type CategoryService struct {
	repo repository.ICategoryRepository
}

func NewCategoryService(repo repository.ICategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (cs *CategoryService) Create(req *schema.CreateCategoryReq) error {
	var insertData model.Category

	insertData.Name = req.Name
	insertData.Description = req.Description

	err := cs.repo.Create(insertData)
	if err != nil {
		return errors.New(reason.CategoryCannotCreate)
	}
	return nil
}

func (cs *CategoryService) BrowseAll() ([]schema.GetCategoryResp, error) {
	var resp []schema.GetCategoryResp

	categories, err := cs.repo.Browse()
	if err != nil {
		return nil, errors.New(reason.CategoryCannotBrowse)
	}

	for _, value := range categories {
		var respData schema.GetCategoryResp
		respData.ID = value.ID
		respData.Name = value.Name
		respData.Description = value.Description
		resp = append(resp, respData)
	}

	return resp, nil
}

func (cs *CategoryService) GetByID(id string) (schema.GetCategoryResp, error) {
	var resp schema.GetCategoryResp

	category, err := cs.repo.GetByID(id)
	if err != nil {
		return resp, errors.New(reason.CategoryCannotGetDetail)
	}

	resp.ID = category.ID
	resp.Name = category.Name
	resp.Description = category.Description

	return resp, nil
}

func (cs *CategoryService) UpdateByID(id string, req *schema.UpdateCategoryReq) error {
	var updateData model.Category

	updateData.Name = req.Name
	updateData.Description = req.Description

	if updateData.ID == 0 {
		return errors.New(reason.CategoryNotFound)
	}

	err := cs.repo.Update(id, updateData)
	if err != nil {
		return errors.New(reason.CategoryCannotUpdate)
	}

	return nil
}

func (cs *CategoryService) DeleteByID(id string) (schema.GetCategoryResp, error) {
	var req schema.GetCategoryResp

	category, err := cs.repo.Delete(id)
	if err != nil {
		return req, errors.New(reason.CategoryCannotDelete)
	}

	req.ID = category.ID
	req.Name = category.Name
	req.Description = category.Description

	if req.ID == 0 {
		return req, errors.New(reason.CategoryNotFound)
	}

	return req, nil
}

/*
func (cs *CategoryService) UpdateID(id string, req schema.UpdateCategoryReq) error {
	// var req schema.UpdateCategoryReq

	category, err := cs.repo.UpdatebyID(id)
	if err != nil {
		return errors.New("cannot update category")
	}

	category.Name = req.Name
	category.Description = req.Description

	return nil
}
func (cs *CategoryService) Updates(id string, req *schema.UpdateCategoryReq) (model.Category, error) {
	// var request schema.UpdateCategoryReq

	category, err := cs.repo.UpdatebyID(id)
	if err != nil {
		return category, errors.New("cannot update category")
	}

	// req.Name = category.Name
	// req.Description = category.Description
	category.Name = req.Name
	category.Description = req.Description

	return category, nil
}

func (cs *CategoryService) CUpdate(resp schema.UpdateCategoryReq) error {
	var insertData model.Category

	insertData.Name = resp.Name
	insertData.Description = resp.Description

	err := cs.repo.CUpdate(insertData)
	if err != nil {
		return errors.New("cannot cupdate category")
	}

	return nil
}
*/
