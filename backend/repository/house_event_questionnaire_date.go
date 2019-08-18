package repository

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/jinzhu/gorm"
)

func HEQDateAll(db *gorm.DB, questionnaireId uint) (*[]model.HouseEventQuestionnaireDate) {
	questionnaire := model.HouseEventQuestionnaire{}

	db.Find(&questionnaire, questionnaireId)
	db.Model(&questionnaire).Related(&questionnaire.Dates, "Dates")

	return &questionnaire.Dates
}

func HEQDateCreate(db *gorm.DB, questionnaireId uint, createQuestionnaireDate *model.HouseEventQuestionnaireDate) *model.HouseEventQuestionnaireDate {
	questionnaire := model.HouseEventQuestionnaire{}
	db.Find(&questionnaire, questionnaireId)
	db.Model(&questionnaire).Association("Dates").Append(createQuestionnaireDate)
	return createQuestionnaireDate
}

func HEQDateFind(db *gorm.DB, questionnaireDateId uint) *model.HouseEventQuestionnaireDate {
	questionnaireDate := model.HouseEventQuestionnaireDate{}
	db.Find(&questionnaireDate, questionnaireDateId).Related(&questionnaireDate.Answers, "Answers")
	return &questionnaireDate
}

func HEQDateUpdate(db *gorm.DB, questionnaireDateId uint, updateQuestionnaireDate *model.HouseEventQuestionnaireDate) *model.HouseEventQuestionnaireDate {
	questionnaireDate := model.HouseEventQuestionnaireDate{}
	db.Find(&questionnaireDate, questionnaireDateId)
	db.Model(&questionnaireDate).Update(updateQuestionnaireDate)
	return &questionnaireDate
}

func HEQDateDelete(db *gorm.DB, questionnaireDateId uint) {
	questionnaireDate := model.HouseEventQuestionnaireDate{}
	db.Delete(&questionnaireDate, questionnaireDateId)
}