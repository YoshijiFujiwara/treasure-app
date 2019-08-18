package validation

import "github.com/YoshijiFujiwara/u22/backend/model"

type HouseEventValidation struct{}

func NewHouseEventValidation() *HouseEventValidation {
	return &HouseEventValidation{}
}

func (hev *HouseEventValidation) CreateValidate(reqParam *model.RequestCreateHouseEvent) bool {
	if reqParam.Name == "" {
		return false
	}
	return true
}

func (hev *HouseEventValidation) UpdateValidation(reqParam *model.RequestUpdateHouseEvent) bool {
	if reqParam.Name == "" {
		return false
	}
	return true
}
