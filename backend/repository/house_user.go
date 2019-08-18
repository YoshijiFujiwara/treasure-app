package repository

import (
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/jinzhu/gorm"
)

func HouseUserFind(db *gorm.DB, userId uint, houseId uint) *model.HouseUser {
	houseUser := model.HouseUser{}
	fmt.Println(userId)
	fmt.Println(houseId)
	fmt.Println("house user find")
	var count int

	db.Debug().Where("user_id = ? AND house_id = ?", userId, houseId).Find(&houseUser).Count(&count)

	fmt.Println("所属ユーザーの情報")
	fmt.Println(houseUser)
	fmt.Println(count)

	return &houseUser
}

func HouseUserIsExists(db *gorm.DB, userId uint, houseId uint) bool {
	houseUser := model.HouseUser{}
	var count int
	db.Debug().Where("user_id = ? AND house_id = ?", userId, houseId).Find(&houseUser).Count(&count)

	if count > 0 {
		return true
	}
	return false
}