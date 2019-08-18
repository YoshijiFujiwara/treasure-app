package model

import "time"

type HouseEventQuestionnaireDate struct {
	BaseModel
	//Questionnaire     HouseEventQuestionnaire             `gorm:"foreignkey:QuestionnaireID" json:"questionnaire"`
	QuestionnaireID   uint                                `json:"questionnaire_id"`
	CandidateDatetime *time.Time                          `json:"candidate_datetime"`
	Answers           []HouseEventQuestionnaireDateAnswer `gorm:"foreignkey:questionnaire_date_id" json: "answers"`
}

type RequestCreateHouseEventQuestionnaireDate struct {
	HouseEventQuestionnaireDate
}

type RequestUpdateHouseEventQuestionnaireDate struct {
	HouseEventQuestionnaireDate
}
