package model

type HouseEventQuestionnaireShop struct {
	BaseModel
	//Questionnaire   House                               `gorm:"foreignkey:QuestionnaireID" json:"questionnaire"`
	QuestionnaireID uint                                `json:"questionnaire_id"`
	ShopID          string                              `json:"shop_id"`
	Answers         []HouseEventQuestionnaireShopAnswer `gorm:"foreignkey:questionnaire_shop_id" json: "answers"`
}

type RequestCreateHouseEventQuestionnaireShop struct {
	HouseEventQuestionnaireShop
}

type RequestUpdateHouseEventQuestionnaireShop struct {
	HouseEventQuestionnaireShop
}
