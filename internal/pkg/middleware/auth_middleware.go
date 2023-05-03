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
		const BearerSchema = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.ErrNoToken)
			ctx.Abort()
			return
		}
		tokenString := authHeader[len(BearerSchema):]

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
