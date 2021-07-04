package models
import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
	Name  string  `gorm:"size:255;index:idx_name,unique" form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	return err
}

func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	return err
}

func GetUser(db *gorm.DB, User *User, id int) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	return err
}

func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

func DeleteUser(db *gorm.DB, User *User, id int) (err error) {
	user := db.Where("id = ?", id).First(User)
	err = user.Error
	user.Delete(User)
	return err
}

