package validation

import "github.com/YoshijiFujiwara/u22/backend/model"

type HouseValidation struct{}

func NewHouseValidation() *HouseValidation {
	return &HouseValidation{}
}

func (hv *HouseValidation) CreateValidate(reqParam *model.RequestCreateHouse) bool {
	if reqParam.Name == "" {
		return false
	}
	return true
}