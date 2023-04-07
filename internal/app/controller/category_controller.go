package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devstore/internal/app/schema"
	"github.com/hafizh24/devstore/internal/app/service"
)

type CategoryController struct {
	service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (cc *CategoryController) BrowseCategory(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	var req schema.CreateCategoryReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = cc.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success create category", "data": req})
}

func (cc *CategoryController) DetailCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	var req schema.UpdateCategoryReq
	postID, _ := ctx.Params.Get("id")

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"status": "error", "message": err.Error()})
		return
	}

	err = cc.service.UpdateByID(postID, req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success update category", "data": &req})
}

func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	_, err := cc.service.DeleteByID(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "data deleted", "id": id})
}

/*
func (cc *CategoryController) Update(ctx *gin.Context) {
	var payload *schema.UpdateCategoryReq
	postID, _ := ctx.Params.Get("id")
	request, err := cc.service.Updates(postID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"service": request, "status": "success", "bind": "kosong", "payload": "kosong"})

	return
}

func (cc *CategoryController) Updates(ctx *gin.Context) {
	postID, _ := ctx.Params.Get("id")
	var req schema.UpdateCategoryReq
	post, err := cc.service.Updates(postID, &req)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	post := cc.service.CUpdate(resp)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"ser": &post, "service": post, "status": "success", "post": &req, "bind": postID, "payload": req})

}
*/
