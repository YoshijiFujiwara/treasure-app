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

type HEQDate struct {
	db *gorm.DB
}

func NewHEQDate(db *gorm.DB) *HEQDate {
	return &HEQDate{db: db}
}

func (heqd *HEQDate) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuesionniareId, ok := vars["questionnaire_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireId, err := strconv.ParseInt(queryQuesionniareId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	heqDateService := service.NewHEQDate(heqd.db)
	questionniareDates, err := heqDateService.All(uint(questionnaireId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, questionniareDates, nil
}

func (heqd *HEQDate) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireId, ok := vars["questionnaire_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireId, err := strconv.ParseInt(queryQuestionnaireId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqParam := &model.RequestCreateHouseEventQuestionnaireDate{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	isValid := validation.NewHEQDateValidation().CreateValidation(reqParam)
	if !isValid {
		return http.StatusBadRequest, nil, nil
	}

	createHouseEventQuestionnaire := &model.HouseEventQuestionnaireDate{
		CandidateDatetime: reqParam.CandidateDatetime,
	}

	heqDateService := service.NewHEQDate(heqd.db)
	questionnaireDate, err := heqDateService.Create(uint(questionnaireId), createHouseEventQuestionnaire)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, questionnaireDate, nil
}

func (heqd *HEQDate) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuesionniareDateId, ok := vars["questionnaire_date_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireDateId, err := strconv.ParseInt(queryQuesionniareDateId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	heqDateService := service.NewHEQDate(heqd.db)
	questionnaireDate, err := heqDateService.FindHEQDateDetail(uint(questionnaireDateId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, questionnaireDate, nil
}

func (heqd *HEQDate) Delete(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuesionniareDateId, ok := vars["questionnaire_date_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireDateId, err := strconv.ParseInt(queryQuesionniareDateId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	heqDateService := service.NewHEQDate(heqd.db)
	err = heqDateService.Delete(uint(questionnaireDateId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusNoContent, nil, nil
}