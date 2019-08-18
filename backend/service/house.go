package service

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/jinzhu/gorm"
)

type House struct {
	db *gorm.DB
}

func NewHouse(db *gorm.DB) *House {
	return &House{db}
}

func (h *House) FindHouseDetail(houseId uint) (*model.House, error) {
	houseDetail := repository.HouseFindWithUser(h.db, houseId)

	return houseDetail, nil
}

func (h *House) Create(createHouse *model.House, user *model.User) (*model.House, error) {
	house := repository.HouseCreate(h.db, createHouse)

	// houseの作成者をAdminとして所属させる
	repository.HouseBelongsUser(h.db, house.ID, user.ID, true)

	return house, nil
}