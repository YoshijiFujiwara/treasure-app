package repository

import (
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/jinzhu/gorm"
)

func HouseAll(db *gorm.DB) (*[]model.House) {
	houses := []model.House{}
	db.Find(&houses)
	return &houses
}

func HouseAllUser(db *gorm.DB, userId uint) (*[]model.House) {
	user := model.User{}
	houses := []model.House{}
	db.First(&user, userId).Related(&houses, "Houses")

	return &houses
}

func HouseFindWithUser(db *gorm.DB, houseId uint) (*model.House) {
	house := model.House{}
	db.First(&house, houseId).Related(&house.Users, "Users").Related(&house.Events, "Events").Related(&house.HouseUser, "HouseUser")

	return &house
}

func HouseCreate(db *gorm.DB, house *model.House) (*model.House) {
	db.Create(&house)
	return house
}

func HouseBelongsUser(db *gorm.DB, houseId uint, userId uint, isAdmin bool) {
	fmt.Println("ハウスを作成します")
	db.Debug().Exec(`
INSERT INTO house_users (house_id, user_id, is_admin) VALUES (?, ?, ?)
`, houseId, userId, isAdmin)
}

func HouseUserUpdateStatus(db *gorm.DB, houseId uint, userId uint, isAdmin bool) {
	db.Exec(`
UPDATE house_users SET is_admin=? WHERE house_id = ? AND user_id = ?
`, isAdmin, houseId, userId)
}

func HouseUserDelete(db *gorm.DB, houseId uint, userId uint) {
	db.Exec(`
DELETE FROM house_users WHERE house_id = ? AND user_id = ?
`, houseId, userId)
}