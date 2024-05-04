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

type SubtraitController struct {
	SubtraitService service.SubtraitService
}

func NewSubtraitController(subtraitService service.SubtraitService) *SubtraitController {
	return &SubtraitController{SubtraitService: subtraitService}
}

func (controller *SubtraitController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	subtraitCreateRequest := request.SubtraitCreateRequest{}
	helper.ReadRequestBody(requests, &subtraitCreateRequest)

	result := controller.SubtraitService.Create(requests.Context(), subtraitCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SubtraitController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	subtraitUpdateRequest := request.SubtraitUpdateRequest{}
	helper.ReadRequestBody(requests, &subtraitUpdateRequest)

	subtraitId := params.ByName("subtraitId")
	id, err := strconv.Atoi(subtraitId)
	helper.PanicIfError(err)
	subtraitUpdateRequest.Id = id
	result := controller.SubtraitService.Update(requests.Context(), subtraitUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SubtraitController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	subtraitId := params.ByName("subtraitId")
	id, err := strconv.Atoi(subtraitId)
	helper.PanicIfError(err)

	result := controller.SubtraitService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SubtraitController) SelectById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	subtraitId := params.ByName("subtraitId")
	id, err := strconv.Atoi(subtraitId)
	helper.PanicIfError(err)

	result := controller.SubtraitService.SelectById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SubtraitController) SelectAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.SubtraitService.SelectAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SubtraitController) SelectByParentId(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	parentId := params.ByName("parentId")
	id, err := strconv.Atoi(parentId)
	helper.PanicIfError(err)

	result := controller.SubtraitService.SelectByParentId(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
