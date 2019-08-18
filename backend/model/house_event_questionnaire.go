package model

type HouseEventQuestionnaire struct {
	BaseModel
	//Event   HouseEvent `gorm:"foreignkey:EventID" json:"event"`
	EventID uint `json:"event_id"`
	Dates   []HouseEventQuestionnaireDate `gorm:"foreignkey:questionnaire_id" json:"dates"`
	Shops   []HouseEventQuestionnaireShop `gorm:"foreignkey:questionnaire_id" json:"shops"`
}

type RequestCreateHouseEventQuestionnaire struct {
	HouseEventQuestionnaire
}

type RequestUpdateHouseEventQuestionnaire struct {
	HouseEventQuestionnaire
}
