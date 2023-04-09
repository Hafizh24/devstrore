package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devstore/internal/app/controller"
	"github.com/hafizh24/devstore/internal/app/repository"
	"github.com/hafizh24/devstore/internal/app/service"
	"github.com/hafizh24/devstore/internal/pkg/config"
	"github.com/hafizh24/devstore/internal/pkg/db"
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
	)
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	// ---------------------------------------------------------------------------------------
	categoryRepository := repository.NewCategoryRepository(DBConn)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories", categoryController.BrowseCategory)
	r.GET("/categories/:id", categoryController.DetailCategory)
	r.PATCH("/categories/:id", categoryController.UpdateCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)
}
