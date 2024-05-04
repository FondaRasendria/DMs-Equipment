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

type TraitController struct {
	TraitService service.TraitService
}

func NewTraitController(traitService service.TraitService) *TraitController {
	return &TraitController{TraitService: traitService}
}

func (controller *TraitController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	traitCreateRequest := request.TraitCreateRequest{}
	helper.ReadRequestBody(requests, &traitCreateRequest)

	result := controller.TraitService.Create(requests.Context(), traitCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *TraitController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	traitUpdateRequest := request.TraitUpdateRequest{}
	helper.ReadRequestBody(requests, &traitUpdateRequest)

	traitId := params.ByName("traitId")
	id, err := strconv.Atoi(traitId)
	helper.PanicIfError(err)
	traitUpdateRequest.Id = id
	result := controller.TraitService.Update(requests.Context(), traitUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *TraitController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	traitId := params.ByName("traitId")
	id, err := strconv.Atoi(traitId)
	helper.PanicIfError(err)

	result := controller.TraitService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *TraitController) SelectById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	traitId := params.ByName("traitId")
	id, err := strconv.Atoi(traitId)
	helper.PanicIfError(err)

	result := controller.TraitService.SelectById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *TraitController) SelectAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.TraitService.SelectAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *TraitController) SelectByParentId(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	parentId := params.ByName("parentId")
	id, err := strconv.Atoi(parentId)
	helper.PanicIfError(err)

	result := controller.TraitService.SelectByParentId(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
