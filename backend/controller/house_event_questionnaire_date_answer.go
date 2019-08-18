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

type HEQDAnswer struct {
	db *gorm.DB
}

func NewHEQDAnswer(db *gorm.DB) *HEQDAnswer {
	return &HEQDAnswer{db: db}
}

func (heqda *HEQDAnswer) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireDateId, ok := vars["questionnaire_date_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireDateId, err := strconv.ParseInt(queryQuestionnaireDateId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	heqdAnswerDateService := service.NewHEQDAnswer(heqda.db)
	dateAnswers, err := heqdAnswerDateService.All(uint(questionnaireDateId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, dateAnswers, nil
}

func (heqda *HEQDAnswer) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireDateId, ok := vars["questionnaire_date_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireDateId, err := strconv.ParseInt(queryQuestionnaireDateId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	reqParam := model.RequestCreateHouseEventQuestionnaireDateAnswer{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	createDateAnswer := &model.HouseEventQuestionnaireDateAnswer{
		QuestionnaireDateID: uint(questionnaireDateId),
		UserID: user.ID,
		Ok: reqParam.Ok,
	}

	heqdaService := service.NewHEQDAnswer(heqda.db)
	answer, err := heqdaService.Create(uint(questionnaireDateId), createDateAnswer)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, answer, nil
}

func (heqda *HEQDAnswer) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireDateAnswerId, ok := vars["questionnaire_date_answer_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireDateAnswerId, err := strconv.ParseInt(queryQuestionnaireDateAnswerId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqParam := model.RequestUpdateHouseEventQuestionnaireDateAnswer{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	updateDateAnswer := &model.HouseEventQuestionnaireDateAnswer{
		Ok: reqParam.Ok,
	}

	dateAnswerService := service.NewHEQDAnswer(heqda.db)
	answer, err := dateAnswerService.Update(uint(questionnaireDateAnswerId), updateDateAnswer)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, answer, nil
}

func (heqda *HEQDAnswer) Delete(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireDateAnswerId, ok := vars["questionnaire_date_answer_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireDateAnswerId, err := strconv.ParseInt(queryQuestionnaireDateAnswerId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	dateAnswerService := service.NewHEQDAnswer(heqda.db)
	err = dateAnswerService.Delete(uint(questionnaireDateAnswerId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusNoContent, nil, nil
}