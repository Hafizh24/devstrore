package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devstore/internal/app/schema"
	"github.com/hafizh24/devstore/internal/app/service"
	"github.com/hafizh24/devstore/internal/pkg/handler"
)

type RegistrationController struct {
	service service.IRegistrationService
}

func NewRegistrationController(service service.IRegistrationService) *RegistrationController {
	return &RegistrationController{service: service}
}

func (rc *RegistrationController) Register(ctx *gin.Context) {
	req := &schema.RegisterReq{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := rc.service.Register(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success register", nil)
}
