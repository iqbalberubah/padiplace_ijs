package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	//product
	GetProducts(c *gin.Context)
	GetProduct(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	//po
	// GetPOS(c *gin.Context)
	GetPO(c *gin.Context)
	CreatePO(c *gin.Context)
	UpdatePO(c *gin.Context)
	DeletePO(c *gin.Context)
}
