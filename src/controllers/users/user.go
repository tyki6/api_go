package userController

import (
	"api_go/src/database"
	"api_go/src/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

func  (repository *UserRepo) Signup(c *gin.Context){
	var user models.User
	errBinding := c.ShouldBind(&user)
	if errBinding != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": errBinding.Error()})
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPassword)
	var mysqlErr mysql.MySQLError
	err := models.CreateUser(repository.Db, &user)
	if err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "duplicate key"})
			return
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user models.User
	errBinding := c.ShouldBind(&user)
	if errBinding != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": errBinding.Error()})
		return
	}
	err := models.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (repository *UserRepo) GetUsers(c *gin.Context) {
	var users []models.User
	err := models.GetUsers(repository.Db, &users)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (repository *UserRepo) GetUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")
	idString, errAtoi := strconv.Atoi(id)
	if errAtoi != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errAtoi.Error()})
		return
	}
	err := models.GetUser(repository.Db, &user, idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (repository *UserRepo) UpdateUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")
	idString, errAtoi := strconv.Atoi(id)
	if errAtoi != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errAtoi.Error()})
		return
	}
	err := models.GetUser(repository.Db, &user, idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	errBinding := c.ShouldBind(&user)
	if errBinding != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"err": errBinding.Error()})
		return
	}

	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepo) DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")
	idString, errAtoi := strconv.Atoi(id)
	if errAtoi != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errAtoi.Error()})
		return
	}
	err := models.DeleteUser(repository.Db, &user, idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
