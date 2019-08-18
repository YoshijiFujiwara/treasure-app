package service

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/jinzhu/gorm"
)

type HEQSAnswer struct {
	db *gorm.DB
}

func NewHEQSAnswer(db *gorm.DB) *HEQSAnswer {
	return &HEQSAnswer{db: db}
}

func (heqsa *HEQSAnswer) All(questionnaireshopId uint) (*[]model.HouseEventQuestionnaireShopAnswer, error) {
	shopAnswers := repository.HEQSAnswerAll(heqsa.db, questionnaireshopId)
	return shopAnswers, nil
}

func (heqsa *HEQSAnswer) Create(questionnaireshopId uint, createShopAnswer *model.HouseEventQuestionnaireShopAnswer) (*model.HouseEventQuestionnaireShopAnswer, error) {
	answer := repository.HEQSAnswerCreate(heqsa.db, questionnaireshopId, createShopAnswer)
	return answer, nil
}

func (heqsa *HEQSAnswer) Update(shopAnswerId uint, shopAnswer *model.HouseEventQuestionnaireShopAnswer) (*model.HouseEventQuestionnaireShopAnswer, error) {
	answer := repository.HEQSAnswerUpdate(heqsa.db, shopAnswerId, shopAnswer)
	return answer, nil
}

func (heqsa *HEQSAnswer) Delete(shopAnswerId uint) error {
	repository.HEQSAnswerDelete(heqsa.db, shopAnswerId)
	return nil
}