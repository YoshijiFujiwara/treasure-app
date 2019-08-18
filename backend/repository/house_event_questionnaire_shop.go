package repository

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/jinzhu/gorm"
)

func HEQShopAll(db *gorm.DB, questionnaireId uint) (*[]model.HouseEventQuestionnaireShop) {
	questionnaire := model.HouseEventQuestionnaire{}

	db.Find(&questionnaire, questionnaireId)
	db.Model(&questionnaire).Related(&questionnaire.Shops, "Shops")

	return &questionnaire.Shops
}

func HEQShopCreate(db *gorm.DB, questionnaireId uint, createQuestionnaireShop *model.HouseEventQuestionnaireShop) *model.HouseEventQuestionnaireShop {
	questionnaire := model.HouseEventQuestionnaire{}
	db.Find(&questionnaire, questionnaireId)
	db.Model(&questionnaire).Association("Shops").Append(createQuestionnaireShop)

	return createQuestionnaireShop
}

func HEQShopFind(db *gorm.DB, questionnaireShopId uint) *model.HouseEventQuestionnaireShop {
	questionnaireShop := model.HouseEventQuestionnaireShop{}
	db.Find(&questionnaireShop, questionnaireShopId).Related(&questionnaireShop.Answers, "Answers")
	return &questionnaireShop
}

func HEQShopupdate(db *gorm.DB, questionnaireShopId uint, updateQuestionnaireShop *model.HouseEventQuestionnaireShop) *model.HouseEventQuestionnaireShop {
	questionnaireShop := model.HouseEventQuestionnaireShop{}
	db.Find(&questionnaireShop, questionnaireShopId)
	db.Model(&questionnaireShop).Update(updateQuestionnaireShop)
	return &questionnaireShop
}

func HEQShopDelete(db *gorm.DB, questionnaireShopId uint) {
	questionnaireShop := model.HouseEventQuestionnaireShop{}
	db.Delete(&questionnaireShop, questionnaireShopId)
}