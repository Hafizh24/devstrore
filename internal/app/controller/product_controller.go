package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devstore/internal/app/schema"
	"github.com/hafizh24/devstore/internal/app/service"
	"github.com/hafizh24/devstore/internal/pkg/handler"
)

type ProductController struct {
	service service.IProductService
}

func NewProductController(service service.IProductService) *ProductController {
	return &ProductController{service: service}
}

func (pc *ProductController) BrowseProduct(ctx *gin.Context) {
	resp, err := pc.service.BrowseAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

func (pc *ProductController) DetailProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := pc.service.GetByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	req := &schema.CreateProductReq{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := pc.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create product", req)
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	_, err := pc.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success deleted product", nil)
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	req := &schema.UpdateProductReq{}
	postID, _ := ctx.Params.Get("id")

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := pc.service.UpdateByID(postID, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update product", nil)
}
