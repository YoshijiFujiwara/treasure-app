package repository

import (
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/jinzhu/gorm"
)

func UserAll(db *gorm.DB) (*[]model.User) {
	users := []model.User{}
	db.Find(&users)
	return &users
}

func UserFind(db *gorm.DB, id uint) (*model.User) {
	user := model.User{}
	db.First(&user)
	return &user
}

func UserCreate(db *gorm.DB, user *model.User) (*model.User) {
	db.Create(&user)
	return user
}

func UserUpdate(db *gorm.DB, id uint, user *model.User) (*model.User) {
	userBefore := model.User{}
	userBefore.ID = id
	db.Model(&userBefore).Update(user)
	fmt.Println("user updated")
	return user
}

func UserDelete(db *gorm.DB, id uint) {
	user := model.User{}
	user.ID = id
	db.Delete(&user)
}

func UserFindByEmail(db *gorm.DB, email string) (*model.User) {
	user := model.User{}
	fmt.Println("email")
	fmt.Println(email)
	db.Debug().Where("email = ?", email).First(&user)
	fmt.Println(user)
	fmt.Println("user info")
	return &user
}
