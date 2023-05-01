package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devstore/internal/app/schema"
	"github.com/hafizh24/devstore/internal/app/service"
	"github.com/hafizh24/devstore/internal/pkg/handler"
)

type RefreshTokenVerifier interface {
	ValidateRefreshToken(tokenString string) (string, error)
}

type SessionController struct {
	service    service.ISessionService
	tokenMaker RefreshTokenVerifier
}

func NewSessionController(service service.ISessionService, tokenMaker RefreshTokenVerifier) *SessionController {
	return &SessionController{service: service, tokenMaker: tokenMaker}
}

func (sc *SessionController) Login(ctx *gin.Context) {
	req := &schema.LoginReq{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	token, err := sc.service.Login(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success login", token)
}
func (sc *SessionController) Refresh(ctx *gin.Context) {
	req := &schema.RefreshTokenReq{}
	refreshToken := ctx.GetHeader("refresh_token")
	if refreshToken == "" {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot refresh token")
		return
	}

	validate, err := sc.tokenMaker.ValidateRefreshToken(refreshToken)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	UserID, _ := strconv.Atoi(validate)

	req.UserID = UserID
	req.RefreshToken = refreshToken

	token, _ := sc.service.RefreshToken(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success refresh token", token)
}

func (sc *SessionController) Logout(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := &schema.LogoutReq{}

	req.UserID = id

	err := sc.service.Logout(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success logout", nil)
}
