package controller

import (
	"encoding/json"
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/controller/validation"
	"github.com/YoshijiFujiwara/u22/backend/httputil"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/service"
	"github.com/jinzhu/gorm"
	"net/http"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) Register(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	fmt.Println("Register")
	reqParam := &model.RequestCreateUser{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	isValid := validation.NewUserValidation().RegisterValidate(reqParam)
	if !isValid {
		return http.StatusBadRequest, nil, nil
	}

	createUser := &model.User{
		Username: reqParam.Username,
		Email:    reqParam.Email,
		Password: reqParam.Password,
		Gender:   reqParam.Gender,
		Birthday: reqParam.Birthday,
	}

	userService := service.NewUser(u.db)
	newUser, err := userService.Create(createUser)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, newUser, nil
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	reqParam := &model.RequestLoginUser{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	isValid := validation.NewUserValidation().LoginValidate(reqParam)
	if !isValid {
		return http.StatusBadRequest, nil, nil
	}

	loginUser := &model.User{
		Email:    reqParam.Email,
		Password: reqParam.Password,
	}

	userService := service.NewUser(u.db)
	jwt, err := userService.Login(loginUser)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, jwt, nil
}

func (u *User) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	reqParam := &model.RequestUpdateUser{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	isValid := validation.NewUserValidation().UpdateValidate(reqParam)
	if !isValid {
		return http.StatusBadRequest, nil, nil
	}

	updateUser := &model.User{
		Username: reqParam.Username,
		Email:    user.Email, // メールアドレスは変わらない前提
		Password: reqParam.Password,
		Gender:   reqParam.Gender,
		Birthday: reqParam.Birthday,
	}

	userService := service.NewUser(u.db)
	updatedUser, err := userService.Update(user.ID, updateUser)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, updatedUser, nil
}

func (u *User) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	userService := service.NewUser(u.db)
	err = userService.Delete(user.ID)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusNoContent, nil, nil
}
