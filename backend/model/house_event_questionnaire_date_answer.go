package model

type HouseEventQuestionnaireDateAnswer struct {
	BaseModel
	//QuestionnaireDate   HouseEventQuestionnaireDate `gorm:"foreignkey:QuestionnaireDateID" json:"questionnaire_date"`
	QuestionnaireDateID uint `json:"questionnaire_date_id"`
	//User                User                        `gorm:"foreignkey:UserID" json:"user"`
	UserID uint `json:"user_id"`
	Ok     bool `json:"ok"`
}

type RequestCreateHouseEventQuestionnaireDateAnswer struct {
	HouseEventQuestionnaireDateAnswer
}

type RequestUpdateHouseEventQuestionnaireDateAnswer struct {
	Ok bool `json:"ok"`
}
