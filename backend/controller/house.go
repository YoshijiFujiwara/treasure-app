package controller

import (
	"encoding/json"
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/controller/validation"
	"github.com/YoshijiFujiwara/u22/backend/httputil"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type House struct {
	db *gorm.DB
}

func NewHouse(db *gorm.DB) *House {
	return &House{db: db}
}

func (h *House) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["house_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	houseId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	fmt.Print(houseId)

	houseService := service.NewHouse(h.db)
	houseDetail, err := houseService.FindHouseDetail(uint(houseId))

	return http.StatusOK, houseDetail, nil
}

func (h *House) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	reqParam := &model.RequestCreateHouse{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	isValid := validation.NewHouseValidation().CreateValidate(reqParam)
	if !isValid {
		return http.StatusBadRequest, nil, nil
	}

	createHouse := &model.House {
		Name: reqParam.Name,
	}
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	houseService := service.NewHouse(h.db)
	house, err := houseService.Create(createHouse, user)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, house, nil
}

