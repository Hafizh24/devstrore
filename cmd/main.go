package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devstore/internal/app/controller"
	"github.com/hafizh24/devstore/internal/app/repository"
	"github.com/hafizh24/devstore/internal/app/service"
	"github.com/hafizh24/devstore/internal/pkg/config"
	"github.com/hafizh24/devstore/internal/pkg/db"
	"github.com/hafizh24/devstore/internal/pkg/handler"
	"github.com/hafizh24/devstore/internal/pkg/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var cfg config.Config
var DBConn *sqlx.DB

func init() {

	configLoad, err := config.LoadConfig(".")
	if err != nil {
		log.Panic("cannot load app config")
	}
	cfg = configLoad

	db, err := db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Panic("db not established")
	}
	DBConn = db

	// Setup logrus
	logLevel, err := log.ParseLevel("debug")
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)                 // appyly log level
	log.SetFormatter(&log.JSONFormatter{}) // define format using json

}

func main() {

	fmt.Println(cfg.DBDriver)
	fmt.Println(cfg.DBConnection)
	fmt.Println(cfg.DBDriver)

	r := gin.New()

	// implement middleware
	r.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
		cors.Default(),
	)

	// ---------------------------------------------------------------------------------------
	tokenMaker := service.NewTokenMaker(
		cfg.AccessTokenKey,
		cfg.RefreshTokenKey,
		cfg.AccessTokenDuration,
		cfg.RefreshTokenDuration,
	)

	categoryRepository := repository.NewCategoryRepository(DBConn)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	productRepository := repository.NewProductRepository(DBConn)
	productService := service.NewProductService(productRepository, categoryRepository)
	productController := controller.NewProductController(productService)

	registrationRepository := repository.NewUserRepository(DBConn)
	registrationService := service.NewRegistrationService(registrationRepository)
	registrationController := controller.NewRegistrationController(registrationService)

	userRepository := repository.NewUserRepository(DBConn)
	authRepository := repository.NewAuthRepository(DBConn)
	sessionService := service.NewSessionService(userRepository, authRepository, tokenMaker)
	sessionController := controller.NewSessionController(sessionService, tokenMaker)

	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories", categoryController.BrowseCategory)
	r.GET("/categories/:id", categoryController.DetailCategory)
	r.PATCH("/categories/:id", categoryController.UpdateCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)

	// product entrypoint
	r.GET("/products", productController.BrowseProduct)
	r.GET("/products/:id", productController.DetailProduct)
	r.POST("/products", productController.CreateProduct)
	r.PATCH("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)

	r.POST("/auth/register", registrationController.Register)

	r.POST("/auth/login", sessionController.Login)
	r.GET("/auth/refresh", sessionController.Refresh)

	secured := r.Group("/secured").Use(middleware.AuthMiddleware(tokenMaker))
	{
		secured.GET("/auth/logout", sessionController.Logout)
		secured.GET("/ping", func(ctx *gin.Context) {
			handler.ResponseSuccess(ctx, http.StatusOK, "pong", nil)
		})
	}

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	// nolint:errcheck
	r.Run(appPort)
}
