package validation

import "github.com/YoshijiFujiwara/u22/backend/model"

type HEQShopValidation struct{}

func NewHEQShopValidation() *HEQShopValidation {
	return &HEQShopValidation{}
}

func (h *HEQShopValidation) CreateValidation(reqParam *model.RequestCreateHouseEventQuestionnaireShop) bool {
	if reqParam.ShopID == "" {
		return false
	}
	return true
}