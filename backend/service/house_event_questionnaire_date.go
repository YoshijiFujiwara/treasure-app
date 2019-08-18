package service

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/jinzhu/gorm"
)

type HEQDate struct {
	db *gorm.DB
}

func NewHEQDate(db *gorm.DB) *HEQDate {
	return &HEQDate{db}
}

func (heqd *HEQDate) All(questionnaireId uint) (*[]model.HouseEventQuestionnaireDate, error) {
	questionnaireDates := repository.HEQDateAll(heqd.db, questionnaireId)
	return questionnaireDates, nil
}

func (heqd *HEQDate) Create(questionniareId uint, createQuestionnaireDate *model.HouseEventQuestionnaireDate) (*model.HouseEventQuestionnaireDate, error) {
	questionnaireDate := repository.HEQDateCreate(heqd.db, questionniareId, createQuestionnaireDate)
	return questionnaireDate, nil
}

func (heqd *HEQDate) FindHEQDateDetail(questionnaireDateId uint) (*model.HouseEventQuestionnaireDate, error) {
	questionnaireDate := repository.HEQDateFind(heqd.db, questionnaireDateId)
	return questionnaireDate, nil
}

func (heqd *HEQDate) Delete(questionnaireDateId uint) error {
	repository.HEQDateDelete(heqd.db, questionnaireDateId)
	return nil
}
