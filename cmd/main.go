package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devstore/internal/app/controller"
	"github.com/hafizh24/devstore/internal/app/repository"
	"github.com/hafizh24/devstore/internal/app/service"
	"github.com/hafizh24/devstore/internal/pkg/config"
	"github.com/hafizh24/devstore/internal/pkg/db"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

}

func main() {

	fmt.Println(cfg.DBDriver)
	fmt.Println(cfg.DBConnection)
	fmt.Println(cfg.DBDriver)

	r := gin.Default()
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

	// r.PUT("/categories/:id", categoryController.UpdateCategory)
	r.PATCH("/categories/:id", categoryController.UpdateCategory)

	r.DELETE("/categories/:id", categoryController.DeleteCategory)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)
}
