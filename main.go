package main

import (
	"fmt"
	"log"

	"padiplace_ijs/app/usecase/crud_po"
	"padiplace_ijs/app/usecase/crud_product"
	"padiplace_ijs/config"
	"padiplace_ijs/handler"
	"padiplace_ijs/pkg/mysql"
	"padiplace_ijs/repository/po_repository"
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

	product_h := handler.ProductHandler(crud_product_uc)

	router.GET("/product", product_h.GetProducts)
	router.GET("/product/:id", product_h.GetProduct)
	router.POST("/product", product_h.CreateProduct)
	router.POST("/product/:id", product_h.UpdateProduct)
	router.DELETE("/product/:id", product_h.DeleteProduct)

	po_repo := po_repository.NewPO(db)
	crud_po_uc := crud_po.NewUseCase(po_repo)

	po_h := handler.POHandler(crud_po_uc)

	// router.GET("/po", po_h.GetPOS)
	router.GET("/po/:id", po_h.GetPO)
	router.POST("/po", po_h.CreatePO)
	router.POST("/po/:id", po_h.UpdatePO)
	router.DELETE("/po/:id", po_h.DeletePO)

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
