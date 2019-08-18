package validation

import "github.com/YoshijiFujiwara/u22/backend/model"

type HEQDateValidation struct{}

func NewHEQDateValidation() *HEQDateValidation {
	return &HEQDateValidation{}
}

func (h *HEQDateValidation) CreateValidation(reqParam *model.RequestCreateHouseEventQuestionnaireDate) bool {
	if reqParam.CandidateDatetime == nil {
		return false
	}
	return true
}