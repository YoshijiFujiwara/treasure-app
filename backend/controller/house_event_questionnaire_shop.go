package controller

import (
	"encoding/json"
	"github.com/YoshijiFujiwara/u22/backend/controller/validation"
	"github.com/YoshijiFujiwara/u22/backend/httputil"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type HEQShop struct {
	db *gorm.DB
}

func NewHEQShop(db *gorm.DB) *HEQShop {
	return &HEQShop{db: db}
}

func (heqd *HEQShop) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireId, ok := vars["questionnaire_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireId, err := strconv.ParseInt(queryQuestionnaireId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	heqShopService := service.NewHEQShop(heqd.db)
	questionnaireShops, err := heqShopService.All(uint(questionnaireId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, questionnaireShops, nil
}

func (heqd *HEQShop) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireId, ok := vars["questionnaire_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireId, err := strconv.ParseInt(queryQuestionnaireId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqParam := &model.RequestCreateHouseEventQuestionnaireShop{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}
	isValid := validation.NewHEQShopValidation().CreateValidation(reqParam)
	if !isValid {
		return http.StatusBadRequest, nil, nil
	}

	createHouseEventQuestionnaire := &model.HouseEventQuestionnaireShop{
		ShopID: reqParam.ShopID,
	}

	heqShopService := service.NewHEQShop(heqd.db)
	questionnaireShop, err := heqShopService.Create(uint(questionnaireId), createHouseEventQuestionnaire)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, questionnaireShop, nil
}

func (heqd *HEQShop) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireShopId, ok := vars["questionnaire_shop_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireShopId, err := strconv.ParseInt(queryQuestionnaireShopId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	heqShopService := service.NewHEQShop(heqd.db)
	questionnaireShop, err := heqShopService.FindHEQShopDetail(uint(questionnaireShopId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, questionnaireShop, nil
}

func (heqd *HEQShop) Delete(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireShopId, ok := vars["questionnaire_shop_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireShopId, err := strconv.ParseInt(queryQuestionnaireShopId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	heqShopService := service.NewHEQShop(heqd.db)
	err = heqShopService.Delete(uint(questionnaireShopId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusNoContent, nil, nil
}
