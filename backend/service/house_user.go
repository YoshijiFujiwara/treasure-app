package service

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/jinzhu/gorm"
)

type HouseUser struct {
	db *gorm.DB
}

func NewHouseUser(db *gorm.DB) *HouseUser {
	return &HouseUser{db}
}

func (hu *HouseUser) Add(addHouseUser *model.HouseUser) *model.House {
	repository.HouseBelongsUser(hu.db, addHouseUser.HouseID, addHouseUser.UserID, addHouseUser.IsAdmin)

	// 更新後のHouse情報でも返すか
	houseDetail := repository.HouseFindWithUser(hu.db, addHouseUser.HouseID)
	return houseDetail
}

func (hu *HouseUser) Update(updateHouseUser *model.HouseUser) *model.House {
	repository.HouseUserUpdateStatus(hu.db, updateHouseUser.HouseID, updateHouseUser.UserID, updateHouseUser.IsAdmin)

	// 更新後のHouse情報でも返すか
	houseDetail := repository.HouseFindWithUser(hu.db, updateHouseUser.HouseID)
	return houseDetail
}

func (hu *HouseUser) Delete(houseId uint, userId uint) *model.House {
	repository.HouseUserDelete(hu.db, houseId, userId)

	// 更新後のHouse情報でも返すか
	houseDetail := repository.HouseFindWithUser(hu.db, houseId)
	return houseDetail
}
