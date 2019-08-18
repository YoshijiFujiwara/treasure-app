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

type HouseEvent struct {
	db *gorm.DB
}

func NewHouseEvent(db *gorm.DB) *HouseEvent {
	return &HouseEvent{db: db}
}

func (he *HouseEvent) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryHouseId, ok := vars["house_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	houseId, err := strconv.ParseInt(queryHouseId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	houseEventService := service.NewHouseEvent(he.db)
	houseEvents, err := houseEventService.All(uint(houseId))

	return http.StatusOK, houseEvents, nil
}

func (he *HouseEvent) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	queryHouseId, ok := vars["house_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	houseId, err := strconv.ParseInt(queryHouseId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqParam := &model.RequestCreateHouseEvent{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	isValid := validation.NewHouseEventValidation().CreateValidate(reqParam)
	if !isValid {
		return http.StatusBadRequest, nil, nil
	}

	createHouseEvent := &model.HouseEvent{
		Name:              reqParam.Name,
		CandidateDatetime: reqParam.CandidateDatetime,
		ShopID:            reqParam.ShopID,
	}

	houseEventService := service.NewHouseEvent(he.db)
	event, err := houseEventService.Create(uint(houseId), createHouseEvent)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, event, nil
}

func (he *HouseEvent) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)

	queryEventId, ok := vars["house_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	eventId, err := strconv.ParseInt(queryEventId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	houseEventService := service.NewHouseEvent(he.db)
	houseEvent, err := houseEventService.FindEventDetail(uint(eventId))

	return http.StatusOK, houseEvent, nil
}

func (he *HouseEvent) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	//queryHouseId, ok := vars["house_id"]
	//if !ok {
	//	return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	//}
	//houseId, err := strconv.ParseInt(queryHouseId, 10, 64)
	//if err != nil {
	//	return http.StatusBadRequest, nil, err
	//}

	queryEventId, ok := vars["event_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	eventId, err := strconv.ParseInt(queryEventId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqParam := &model.RequestUpdateHouseEvent{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	isValid := validation.NewHouseEventValidation().UpdateValidation(reqParam)
	if !isValid {
		return http.StatusBadRequest, nil, nil
	}

	updateHouseEvent := &model.HouseEvent{
		Name:              reqParam.Name,
		CandidateDatetime: reqParam.CandidateDatetime,
		ShopID:            reqParam.ShopID,
	}

	houseEventService := service.NewHouseEvent(he.db)
	event, err := houseEventService.Update(uint(eventId), updateHouseEvent)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, event, nil
}

func (he *HouseEvent) Delete(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	//queryHouseId, ok := vars["house_id"]
	//if !ok {
	//	return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	//}
	//houseId, err := strconv.ParseInt(queryHouseId, 10, 64)
	//if err != nil {
	//	return http.StatusBadRequest, nil, err
	//}

	queryEventId, ok := vars["event_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	eventId, err := strconv.ParseInt(queryEventId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	houseEventService := service.NewHouseEvent(he.db)
	err = houseEventService.Delete(uint(eventId))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusNoContent, nil, nil
}

