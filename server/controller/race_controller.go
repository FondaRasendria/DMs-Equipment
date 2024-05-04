package controller

import (
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type RaceController struct {
	RaceService service.RaceService
}

func NewRaceController(raceService service.RaceService) *RaceController {
	return &RaceController{RaceService: raceService}
}

func (controller *RaceController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	raceCreateRequest := request.RaceCreateRequest{}
	helper.ReadRequestBody(requests, &raceCreateRequest)

	result := controller.RaceService.Create(requests.Context(), raceCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *RaceController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	raceUpdateRequest := request.RaceUpdateRequest{}
	helper.ReadRequestBody(requests, &raceUpdateRequest)

	raceId := params.ByName("raceId")
	id, err := strconv.Atoi(raceId)
	helper.PanicIfError(err)
	raceUpdateRequest.Id = id
	result := controller.RaceService.Update(requests.Context(), raceUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *RaceController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	raceId := params.ByName("raceId")
	id, err := strconv.Atoi(raceId)
	helper.PanicIfError(err)

	result := controller.RaceService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *RaceController) SelectById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	raceId := params.ByName("raceId")
	id, err := strconv.Atoi(raceId)
	helper.PanicIfError(err)

	result := controller.RaceService.SelectById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *RaceController) SelectAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.RaceService.SelectAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *RaceController) SelectByName(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	name := params.ByName("name")

	result := controller.RaceService.SelectByName(requests.Context(), name)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
