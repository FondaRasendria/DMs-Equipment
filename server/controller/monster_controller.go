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

type MonsterController struct {
	MonsterService service.MonsterService
}

func NewMonsterController(monsterService service.MonsterService) *MonsterController {
	return &MonsterController{MonsterService: monsterService}
}

func (controller *MonsterController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	monsterCreateRequest := request.MonsterCreateRequest{}
	helper.ReadRequestBody(requests, &monsterCreateRequest)

	result := controller.MonsterService.Create(requests.Context(), monsterCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	monsterUpdateRequest := request.MonsterUpdateRequest{}
	helper.ReadRequestBody(requests, &monsterUpdateRequest)

	monsterId := params.ByName("monsterId")
	id, err := strconv.Atoi(monsterId)
	helper.PanicIfError(err)
	monsterUpdateRequest.Id = id
	result := controller.MonsterService.Update(requests.Context(), monsterUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	monsterId := params.ByName("monsterId")
	id, err := strconv.Atoi(monsterId)
	helper.PanicIfError(err)

	result := controller.MonsterService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterController) SelectById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	monsterId := params.ByName("monsterId")
	id, err := strconv.Atoi(monsterId)
	helper.PanicIfError(err)

	result := controller.MonsterService.SelectById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterController) SelectAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.MonsterService.SelectAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterController) SelectByName(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	name := params.ByName("name")

	result := controller.MonsterService.SelectByName(requests.Context(), name)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterController) SelectBySource(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	source := params.ByName("source")

	result := controller.MonsterService.SelectBySource(requests.Context(), source)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
