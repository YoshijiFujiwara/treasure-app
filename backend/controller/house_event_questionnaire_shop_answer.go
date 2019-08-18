package controller

import (
	"encoding/json"
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/httputil"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type HEQSAnswer struct {
	db *gorm.DB
}

func NewHEQSAnswer(db *gorm.DB) *HEQSAnswer {
	return &HEQSAnswer{db: db}
}

func (heqsa *HEQSAnswer) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireShopId, ok := vars["questionnaire_shop_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireShopId, err := strconv.ParseInt(queryQuestionnaireShopId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	heqdAnswerShopService := service.NewHEQSAnswer(heqsa.db)
	shopAnswers, err := heqdAnswerShopService.All(uint(questionnaireShopId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, shopAnswers, nil
}

func (heqsa *HEQSAnswer) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireShopId, ok := vars["questionnaire_shop_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireShopId, err := strconv.ParseInt(queryQuestionnaireShopId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	reqParam := model.RequestCreateHouseEventQuestionnaireShopAnswer{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	createShopAnswer := &model.HouseEventQuestionnaireShopAnswer{
		QuestionnaireShopID: uint(questionnaireShopId),
		UserID: user.ID,
		Ok: reqParam.Ok,
	}

	fmt.Println("createShopAnswer")
	fmt.Println(createShopAnswer)

	heqsaService := service.NewHEQSAnswer(heqsa.db)
	answer, err := heqsaService.Create(uint(questionnaireShopId), createShopAnswer)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, answer, nil
}

func (heqsa *HEQSAnswer) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireShopAnswerId, ok := vars["questionnaire_shop_answer_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireShopAnswerId, err := strconv.ParseInt(queryQuestionnaireShopAnswerId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqParam := model.RequestUpdateHouseEventQuestionnaireShopAnswer{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	updateShopAnswer := &model.HouseEventQuestionnaireShopAnswer{
		Ok: reqParam.Ok,
	}

	shopAnswerService := service.NewHEQSAnswer(heqsa.db)
	answer, err := shopAnswerService.Update(uint(questionnaireShopAnswerId), updateShopAnswer)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, answer, nil
}

func (heqsa *HEQSAnswer) Delete(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryQuestionnaireShopAnswerId, ok := vars["questionnaire_shop_answer_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	questionnaireShopAnswerId, err := strconv.ParseInt(queryQuestionnaireShopAnswerId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	shopAnswerService := service.NewHEQSAnswer(heqsa.db)
	err = shopAnswerService.Delete(uint(questionnaireShopAnswerId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusNoContent, nil, nil
}
