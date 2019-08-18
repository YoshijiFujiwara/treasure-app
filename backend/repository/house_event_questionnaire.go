package repository

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/jinzhu/gorm"
)

func HouseEventQuestionnaireAll(db *gorm.DB, eventId uint) (*[]model.HouseEventQuestionnaire) {
	event := model.HouseEvent{}

	db.Find(&event, eventId)
	db.Debug().Model(&event).Related(&event.Questionnaires, "Questionnaires")

	return &event.Questionnaires
}

func HouseEventQuestionnaireCreate(db *gorm.DB, eventId uint, createQuestionnaire *model.HouseEventQuestionnaire) *model.HouseEventQuestionnaire {
	event := model.HouseEvent{}
	db.Find(&event, eventId)
	db.Model(&event).Association("Questionnaires").Append(createQuestionnaire)

	return createQuestionnaire
}

func HouseEventQuestionnaireFind(db *gorm.DB, questionnaireId uint) *model.HouseEventQuestionnaire {
	questionnaire := model.HouseEventQuestionnaire{}
	db.Find(&questionnaire, questionnaireId).Related(&questionnaire.Dates, "Dates").Related(&questionnaire.Shops, "Shops")
	return &questionnaire
}

func HouseEventQuestionnaireDelete(db *gorm.DB, questionnaireId uint) {
	questionnaire := model.HouseEventQuestionnaire{}
	db.Delete(&questionnaire, questionnaireId)
}