package validation

import "github.com/YoshijiFujiwara/u22/backend/model"

type UserValidation struct{}

func NewUserValidation() *UserValidation {
	return &UserValidation{}
}

func (uv *UserValidation) RegisterValidate(reqParam *model.RequestCreateUser) bool {
	if reqParam.Email == "" || reqParam.Password == "" || reqParam.Username == "" {
		return false
	}
	return true
}

func (uv *UserValidation) LoginValidate(reqParam *model.RequestLoginUser) bool {
	if reqParam.Password == "" || reqParam.Email == "" {
		return false
	}
	return true
}

func (uv *UserValidation) UpdateValidate(reqParam *model.RequestUpdateUser) bool {
	if reqParam.Username == "" || reqParam.Password == "" {
		return false
	}
	return true
}
