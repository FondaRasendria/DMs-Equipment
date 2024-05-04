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

type ActionController struct {
	ActionService service.ActionService
}

func NewActionController(actionService service.ActionService) *ActionController {
	return &ActionController{ActionService: actionService}
}

func (controller *ActionController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	actionCreateRequest := request.ActionCreateRequest{}
	helper.ReadRequestBody(requests, &actionCreateRequest)

	result := controller.ActionService.Create(requests.Context(), actionCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ActionController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	actionUpdateRequest := request.ActionUpdateRequest{}
	helper.ReadRequestBody(requests, &actionUpdateRequest)

	actionId := params.ByName("actionId")
	id, err := strconv.Atoi(actionId)
	helper.PanicIfError(err)
	actionUpdateRequest.Id = id
	result := controller.ActionService.Update(requests.Context(), actionUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ActionController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	actionId := params.ByName("actionId")
	id, err := strconv.Atoi(actionId)
	helper.PanicIfError(err)

	result := controller.ActionService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ActionController) SelectById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	actionId := params.ByName("actionId")
	id, err := strconv.Atoi(actionId)
	helper.PanicIfError(err)

	result := controller.ActionService.SelectById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ActionController) SelectAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.ActionService.SelectAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ActionController) SelectByParentId(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	parentId := params.ByName("parentId")
	id, err := strconv.Atoi(parentId)
	helper.PanicIfError(err)

	result := controller.ActionService.SelectByParentId(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
