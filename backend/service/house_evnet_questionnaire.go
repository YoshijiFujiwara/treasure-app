package service

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/jinzhu/gorm"
)

type HouseEventQuestionnaire struct {
	db *gorm.DB
}

func NewHouseEventQuestionnaire(db *gorm.DB) *HouseEventQuestionnaire {
	return &HouseEventQuestionnaire{db}
}

func (heq *HouseEventQuestionnaire) All(eventId uint) (*[]model.HouseEventQuestionnaire, error) {
	houseEventQuestionnaires := repository.HouseEventQuestionnaireAll(heq.db, eventId)
	return houseEventQuestionnaires, nil
}

func (heq *HouseEventQuestionnaire) Create(eventId uint, createQuestionnaire *model.HouseEventQuestionnaire) (*model.HouseEventQuestionnaire, error) {
	newQuestionnaire := repository.HouseEventQuestionnaireCreate(heq.db, eventId, createQuestionnaire)
	return newQuestionnaire, nil
}

func (heq *HouseEventQuestionnaire) FindQuestionnaireDetail(questionnaireId uint) (*model.HouseEventQuestionnaire, error) {
	questionnaire := repository.HouseEventQuestionnaireFind(heq.db, questionnaireId)
	return questionnaire, nil
}

func (heq *HouseEventQuestionnaire) Delete(questionnaireId uint) error {
	repository.HouseEventQuestionnaireDelete(heq.db, questionnaireId)
	return nil
}