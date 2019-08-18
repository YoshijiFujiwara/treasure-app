package service

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/jinzhu/gorm"
)

type HEQShop struct {
	db *gorm.DB
}

func NewHEQShop(db *gorm.DB) *HEQShop {
	return &HEQShop{db}
}

func (heqd *HEQShop) All(questionnaireId uint) (*[]model.HouseEventQuestionnaireShop, error) {
	questionnaireShop := repository.HEQShopAll(heqd.db, questionnaireId)
	return questionnaireShop, nil
}

func (heqd *HEQShop) Create(questionniareId uint, createQuestionnaireShop *model.HouseEventQuestionnaireShop) (*model.HouseEventQuestionnaireShop, error) {
	newQuestionnaireShop := repository.HEQShopCreate(heqd.db, questionniareId, createQuestionnaireShop)
	return newQuestionnaireShop, nil
}

func (heqd *HEQShop) FindHEQShopDetail(questionnaireShopId uint) (*model.HouseEventQuestionnaireShop, error) {
	questionnaireShop := repository.HEQShopFind(heqd.db, questionnaireShopId)
	return questionnaireShop, nil
}

func (heqd *HEQShop) Delete(questionnaireShopId uint) error {
	repository.HEQShopDelete(heqd.db, questionnaireShopId)
	return nil
}
