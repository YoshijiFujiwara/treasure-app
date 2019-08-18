package model

import "time"

type HouseEvent struct {
	BaseModel
	//House             House                     `gorm:"foreignkey:HouseID" json:"house"`
	HouseID           uint                      `json:"house_id"`
	Name              string                    `json:"name"`
	CandidateDatetime *time.Time                `json:"candidate_datetime"`
	ShopID            string                    `json:"shop_id"`
	Questionnaires    []HouseEventQuestionnaire `gorm:"foreignkey:event_id" json:"questionnaires"`
}

type RequestCreateHouseEvent struct {
	Name              string     `json:"name"`
	CandidateDatetime *time.Time `json:"candidate_datetime"`
	ShopID            string     `json:"shop_id"`
}

type RequestUpdateHouseEvent struct {
	Name              string     `json:"name"`
	CandidateDatetime *time.Time `json:"candidate_datetime"`
	ShopID            string     `json:"shop_id"`
}
