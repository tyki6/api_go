package main

import (
	"api_go/src/controllers/products"
	"api_go/src/controllers/users"
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
	//product routes
	r.GET("/products", productRepo.GetProducts)
	r.GET("/products/:id", productRepo.GetProduct)
	r.POST("/products", productRepo.CreateProduct)
	r.PUT("/products/:id", productRepo.UpdateProduct)
	r.DELETE("/products/:id", productRepo.DeleteProduct)
	userRepo := userController.New()
	//user routes
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.POST("/users", userRepo.CreateUser)
	r.POST("/signup", userRepo.Signup)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)
	r.Run() // listen and serve on 0.0.0.0:8080
}
