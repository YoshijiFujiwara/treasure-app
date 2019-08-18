package controller

import (
	"encoding/json"
	"github.com/YoshijiFujiwara/u22/backend/httputil"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type HouseUser struct {
	db *gorm.DB
}

func NewHouseUser(db *gorm.DB) *HouseUser {
	return &HouseUser{db: db}
}


func (hu *HouseUser) Add(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryHouseId, ok := vars["house_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	houseId, err := strconv.ParseInt(queryHouseId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	queryUserId, ok := vars["user_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	userId, err := strconv.ParseInt(queryUserId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqParam := &model.RequestAddHouseUser{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	addHouseUser := &model.HouseUser{
		UserID: uint(userId),
		HouseID: uint(houseId),
		IsAdmin: reqParam.IsAdmin,
	}

	houseUserService := service.NewHouseUser(hu.db)
	houseDetail := houseUserService.Add(addHouseUser)

	return http.StatusOK, houseDetail, nil
}

func (hu *HouseUser) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryHouseId, ok := vars["house_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	houseId, err := strconv.ParseInt(queryHouseId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	queryUserId, ok := vars["user_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	userId, err := strconv.ParseInt(queryUserId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqParam := &model.RequestUpdateHouseUser{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	updateHouseUser := &model.HouseUser{
		UserID: uint(userId),
		HouseID: uint(houseId),
		IsAdmin: reqParam.IsAdmin,
	}

	houseUserService := service.NewHouseUser(hu.db)
	houseDetail := houseUserService.Update(updateHouseUser)

	return http.StatusOK, houseDetail, nil
}


func (hu *HouseUser) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryHouseId, ok := vars["house_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	houseId, err := strconv.ParseInt(queryHouseId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	queryUserId, ok := vars["user_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	userId, err := strconv.ParseInt(queryUserId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	houseUserService := service.NewHouseUser(hu.db)
	houseDetail := houseUserService.Delete(uint(houseId), uint(userId))
	return http.StatusOK, houseDetail, nil
}