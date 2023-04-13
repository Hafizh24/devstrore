package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoggingMiddleware(t *testing.T) {

	type TestCase struct {
		Name            string
		ResponseCode    int
		ResponseMessage string
		FuncHandler     func(c *gin.Context)
	}

	cases := []TestCase{
		{
			Name:            "request success",
			ResponseCode:    http.StatusOK,
			ResponseMessage: "success response",
			FuncHandler: func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success response"})
			},
		},
		{
			Name:            "request error",
			ResponseCode:    http.StatusUnprocessableEntity,
			ResponseMessage: "success response",
			FuncHandler: func(c *gin.Context) {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error"})
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.Use(LoggingMiddleware())
			router.GET("/test", tc.FuncHandler)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/test", nil)
			router.ServeHTTP(response, request)

			// CHECK ASSERT
			assert.Equal(t, tc.ResponseCode, response.Code)

		})
	}
}
