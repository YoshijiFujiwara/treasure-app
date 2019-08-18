package model

type HouseEventQuestionnaireShopAnswer struct {
	BaseModel
	//QuestionnaireShop   HouseEventQuestionnaireShop `gorm:"foreignkey:QuestionnaireShopID" json:"questionnaire_shop"`
	QuestionnaireShopID uint `json:"questionnaire_shop_id"`
	User                User `gorm:"foreignkey:UserID" json:"user"`
	UserID              uint `json: "user_id"`
	Ok                  bool `json: "ok"`
}

type RequestCreateHouseEventQuestionnaireShopAnswer struct {
	HouseEventQuestionnaireShopAnswer
}

type RequestUpdateHouseEventQuestionnaireShopAnswer struct {
	Ok bool `json: "ok"`
}
