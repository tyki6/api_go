package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id    int     `gorm:"primaryKey" form:"id" json:"id"`
	Name  string  `form:"name" json:"name" binding:"required"`
	Price float32 `form:"price" json:"price" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateProduct(db *gorm.DB, Product *Product) (err error) {
	err = db.Create(Product).Error
	return err
}

func GetProducts(db *gorm.DB, Product *[]Product) (err error) {
	err = db.Find(Product).Error
	return err
}

func GetProduct(db *gorm.DB, Product *Product, id int) (err error) {
	err = db.Where("id = ?", id).First(Product).Error
	return err
}

func UpdateProduct(db *gorm.DB, Product *Product) (err error) {
	db.Save(Product)
	return nil
}

func DeleteProduct(db *gorm.DB, Product *Product, id int) (err error) {
	product := db.Where("id = ?", id).First(Product)
	err = product.Error
	product.Delete(Product)
	return err
}
