package main

import (
	"api_go/src/controllers/products"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	productRepo := productController.New()
	r.GET("/products", productRepo.GetProducts)
	r.GET("/products/:id", productRepo.GetProduct)
	r.POST("/products", productRepo.CreateProduct)
	r.PUT("/products/:id", productRepo.UpdateProduct)
	r.DELETE("/products/:id", productRepo.DeleteProduct)
	r.Run() // listen and serve on 0.0.0.0:8080
}
