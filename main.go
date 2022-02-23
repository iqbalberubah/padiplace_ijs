package main

import (
	"fmt"
	"log"

	"padiplace_ijs/app/usecase/crud_product"
	"padiplace_ijs/config"
	"padiplace_ijs/handler"
	"padiplace_ijs/pkg/mysql"
	"padiplace_ijs/repository/product_repository"

	"gorm.io/gorm/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.New()
	router := gin.Default()

	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName, logger.LogLevel(cfg.DBLogMode))
	if err != nil {
		log.Panicln("Failed to Initialized mysql DB:", err)
	}

	product_repo := product_repository.NewProduct(db)
	crud_product_uc := crud_product.NewUseCase(product_repo)

	h := handler.NewHandler(crud_product_uc)

	router.GET("/product", h.GetProducts)
	router.GET("/product/:id", h.GetProduct)
	router.POST("/product", h.CreateProduct)
	router.POST("/product/:id", h.UpdateProduct)
	router.DELETE("/product/:id", h.DeleteProduct)

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
