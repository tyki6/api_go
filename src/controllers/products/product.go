package productController

import (
	"api_go/src/database"
	"api_go/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ProductRepo struct {
	Db *gorm.DB
}

func New() *ProductRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Product{})
	return &ProductRepo{Db: db}
}

func (repository *ProductRepo) CreateProduct(c *gin.Context) {
	var product models.Product
	errBinding := c.ShouldBind(&product)
	if errBinding != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": errBinding.Error()})
		return
	}
	err := models.CreateProduct(repository.Db, &product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (repository *ProductRepo) GetProducts(c *gin.Context) {
	var products []models.Product
	err := models.GetProducts(repository.Db, &products)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (repository *ProductRepo) GetProduct(c *gin.Context) {
	var product models.Product
	id, _ := c.Params.Get("id")
	idString, errAtoi := strconv.Atoi(id)
	if errAtoi != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errAtoi.Error()})
		return
	}
	err := models.GetProduct(repository.Db, &product, idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (repository *ProductRepo) UpdateProduct(c *gin.Context) {
	var product models.Product
	id, _ := c.Params.Get("id")
	idString, errAtoi := strconv.Atoi(id)
	if errAtoi != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errAtoi.Error()})
		return
	}
	err := models.GetProduct(repository.Db, &product, idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	errBinding := c.ShouldBind(&product)
	if errBinding != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"err": errBinding.Error()})
		return
	}

	err = models.UpdateProduct(repository.Db, &product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (repository *ProductRepo) DeleteProduct(c *gin.Context) {
	var product models.Product
	id, _ := c.Params.Get("id")
	idString, errAtoi := strconv.Atoi(id)
	if errAtoi != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errAtoi.Error()})
		return
	}
	err := models.DeleteProduct(repository.Db, &product, idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
