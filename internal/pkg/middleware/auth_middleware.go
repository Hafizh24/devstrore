package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hafizh24/devstore/internal/pkg/handler"
	"github.com/hafizh24/devstore/internal/pkg/reason"
)

type AccessTokenVerifier interface {
	ValidateAccessToken(tokenString string) (string, error)
}

func AuthMiddleware(tokenMaker AccessTokenVerifier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.ErrNoToken)
			ctx.Abort()
			return
		}
		sub, err := tokenMaker.ValidateAccessToken(tokenString)
		if err != nil {
			handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		ctx.Set("user_id", sub)
		ctx.Next()
	}
}
