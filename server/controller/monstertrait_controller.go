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

type MonsterTraitController struct {
	MonsterTraitService service.MonsterTraitService
}

func NewMonsterTraitController(mtService service.MonsterTraitService) *MonsterTraitController {
	return &MonsterTraitController{MonsterTraitService: mtService}
}

func (controller *MonsterTraitController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	mtCreateRequest := request.MonsterTraitCreateRequest{}
	helper.ReadRequestBody(requests, &mtCreateRequest)

	result := controller.MonsterTraitService.Create(requests.Context(), mtCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterTraitController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	mtUpdateRequest := request.MonsterTraitUpdateRequest{}
	helper.ReadRequestBody(requests, &mtUpdateRequest)

	mtId := params.ByName("mtId")
	id, err := strconv.Atoi(mtId)
	helper.PanicIfError(err)
	mtUpdateRequest.Id = id
	result := controller.MonsterTraitService.Update(requests.Context(), mtUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterTraitController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	mtId := params.ByName("mtId")
	id, err := strconv.Atoi(mtId)
	helper.PanicIfError(err)

	result := controller.MonsterTraitService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterTraitController) SelectById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	mtId := params.ByName("mtId")
	id, err := strconv.Atoi(mtId)
	helper.PanicIfError(err)

	result := controller.MonsterTraitService.SelectById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterTraitController) SelectAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.MonsterTraitService.SelectAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *MonsterTraitController) SelectByParentId(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	parentId := params.ByName("parentId")
	id, err := strconv.Atoi(parentId)
	helper.PanicIfError(err)

	result := controller.MonsterTraitService.SelectByParentId(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
