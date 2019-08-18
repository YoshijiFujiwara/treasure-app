package service

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/jinzhu/gorm"
)

type HouseEvent struct {
	db *gorm.DB
}

func NewHouseEvent(db *gorm.DB) *HouseEvent {
	return &HouseEvent{db}
}


func (he *HouseEvent) All(houseId uint) (*[]model.HouseEvent, error) {
	houseEvents := repository.HouseEventAll(he.db, houseId)
	return houseEvents, nil
}

func (he *HouseEvent) Create(houseId uint, createHouseEvent *model.HouseEvent) (*model.HouseEvent, error) {
	event := repository.HouseEventCreate(he.db, houseId, createHouseEvent)
	return event, nil
}

func (he *HouseEvent) FindEventDetail(eventId uint) (*model.HouseEvent, error) {
	event := repository.HouseEventFind(he.db, eventId)
	return event, nil
}

func (he *HouseEvent) Update(eventId uint, updateHouseEvent *model.HouseEvent) (*model.HouseEvent, error) {
	repository.HouseEventUpdate(he.db, eventId, updateHouseEvent)

	// イベントの詳細
	eventDetail := repository.HouseEventFind(he.db, eventId)
	return eventDetail, nil
}

func (he *HouseEvent) Delete(eventId uint) (error) {
	repository.HouseEventDelete(he.db, eventId)
	return nil
}