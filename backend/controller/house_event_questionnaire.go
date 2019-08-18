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

type HouseEventQuestionnaire struct {
	db *gorm.DB
}

func NewHouseEventQuestionnaire(db *gorm.DB) *HouseEventQuestionnaire {
	return &HouseEventQuestionnaire{db: db}
}

func (heq *HouseEventQuestionnaire) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryEventId, ok := vars["event_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	eventId, err := strconv.ParseInt(queryEventId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	houseEventQuestionnaireService := service.NewHouseEventQuestionnaire(heq.db)
	questionnaires, err := houseEventQuestionnaireService.All(uint(eventId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, questionnaires, nil
}

func (heq *HouseEventQuestionnaire) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryEventId, ok := vars["event_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	eventId, err := strconv.ParseInt(queryEventId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqParam := model.RequestCreateHouseEventQuestionnaire{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	createHouseEventQuestionnaire := &model.HouseEventQuestionnaire{}

	houseEventQuestionnaireService := service.NewHouseEventQuestionnaire(heq.db)
	newQuestionnaire, err := houseEventQuestionnaireService.Create(uint(eventId), createHouseEventQuestionnaire)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, newQuestionnaire, nil
}

func (heq *HouseEventQuestionnaire) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireId, ok := vars["questionnaire_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireId, err := strconv.ParseInt(queryQuestionnaireId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	houseEventQuestionnaireService := service.NewHouseEventQuestionnaire(heq.db)
	questionnaire, err := houseEventQuestionnaireService.FindQuestionnaireDetail(uint(questionnaireId))

	return http.StatusOK, questionnaire, nil
}

func (heq *HouseEventQuestionnaire) Delete(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)

	queryQuestionnaireId, ok := vars["questionnaire_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireId, err := strconv.ParseInt(queryQuestionnaireId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	houseEventQuestionnaireService := service.NewHouseEventQuestionnaire(heq.db)
	err = houseEventQuestionnaireService.Delete(uint(questionnaireId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusNoContent, nil, nil
}