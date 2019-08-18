package service

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/jinzhu/gorm"
)

type HEQDAnswer struct {
	db *gorm.DB
}

func NewHEQDAnswer(db *gorm.DB) *HEQDAnswer {
	return &HEQDAnswer{db: db}
}

func (heqda *HEQDAnswer) All(questionnaireDateId uint) (*[]model.HouseEventQuestionnaireDateAnswer, error) {
	dateAnswers := repository.HEQDAnswerAll(heqda.db, questionnaireDateId)
	return dateAnswers, nil
}

func (heqda *HEQDAnswer) Create(questionnaireDateId uint, createDateAnswer *model.HouseEventQuestionnaireDateAnswer) (*model.HouseEventQuestionnaireDateAnswer, error) {
	answer := repository.HEQDAnswerCreate(heqda.db, questionnaireDateId, createDateAnswer)
	return answer, nil
}

func (heqda *HEQDAnswer) Update(dateAnswerId uint, updateDateAnswer *model.HouseEventQuestionnaireDateAnswer) (*model.HouseEventQuestionnaireDateAnswer, error) {
	answer := repository.HEQDAnswerUpdate(heqda.db, dateAnswerId, updateDateAnswer)
	return answer, nil
}

func (heqda *HEQDAnswer) Delete(dateAnswerId uint) error {
	repository.HEQDAnswerDelete(heqda.db, dateAnswerId)
	return nil
}