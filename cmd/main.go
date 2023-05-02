package main

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
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

var (
	cfg      config.Config
	DBConn   *sqlx.DB
	enforcer *casbin.Enforcer
)

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

	// setup casbin
	e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic("cannot load app casbin enforcer")
	}
	enforcer = e

}

func main() {

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

	r.POST("/auth/register", registrationController.Register)

	r.POST("/auth/login", sessionController.Login)
	r.GET("/auth/refresh", sessionController.Refresh)

	secured := r.Group("/api").Use(middleware.AuthMiddleware(tokenMaker))
	{
		secured.GET("/auth/logout", sessionController.Logout)
		secured.GET("/ping", func(ctx *gin.Context) {
			handler.ResponseSuccess(ctx, http.StatusOK, "pong", nil)
		})

		secured.GET("/categories", middleware.AuthorizationMiddleware("categories", "read", enforcer), categoryController.BrowseCategory)
		secured.GET("/categories/:id", middleware.AuthorizationMiddleware("categories", "read", enforcer), categoryController.DetailCategory)
		secured.POST("/categories", middleware.AuthorizationMiddleware("categories", "write", enforcer), categoryController.CreateCategory)
		secured.PATCH("/categories/:id", middleware.AuthorizationMiddleware("categories", "write", enforcer), categoryController.UpdateCategory)
		secured.DELETE("/categories/:id", middleware.AuthorizationMiddleware("categories", "write", enforcer), categoryController.DeleteCategory)

		secured.GET("/products", middleware.AuthorizationMiddleware("products", "read", enforcer), productController.BrowseProduct)
		secured.GET("/products/:id", middleware.AuthorizationMiddleware("products", "read", enforcer), productController.DetailProduct)
		secured.POST("/products", middleware.AuthorizationMiddleware("products", "write", enforcer), productController.CreateProduct)
		secured.PATCH("/products/:id", middleware.AuthorizationMiddleware("products", "write", enforcer), productController.UpdateProduct)
		secured.DELETE("/products/:id", middleware.AuthorizationMiddleware("products", "write", enforcer), productController.DeleteProduct)
	}

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	// nolint:errcheck
	r.Run(appPort)
}
