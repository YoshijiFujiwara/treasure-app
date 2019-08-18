package repository

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/jinzhu/gorm"
)

func HouseEventAll(db *gorm.DB, houseId uint) (*[]model.HouseEvent) {
	house := model.House{}
	houseEvents := []model.HouseEvent{}

	db.Find(&house, houseId)
	db.Model(&house).Related(&houseEvents)

	return &houseEvents
}

func HouseEventCreate(db *gorm.DB, houseId uint, createHouseEvent *model.HouseEvent) *model.HouseEvent {
	house := model.House{}
	db.Find(&house, houseId)
	db.Model(&house).Association("Events").Append(createHouseEvent)
	return createHouseEvent
}

func HouseEventFind(db *gorm.DB, eventId uint) *model.HouseEvent {
	event := model.HouseEvent{}

	db.Find(&event, eventId)
	db.Model(&event).Related(&event.Questionnaires, "Questionnaires")
	return &event
}

func HouseEventUpdate(db *gorm.DB, eventId uint, updateHouseEvent *model.HouseEvent) *model.HouseEvent {
	event := model.HouseEvent{}
	db.Find(&event, eventId)
	db.Model(&event).Update(updateHouseEvent)

	return &event
}

func HouseEventDelete(db *gorm.DB, eventId uint) {
	event := model.HouseEvent{}
	db.Delete(&event, eventId)
}