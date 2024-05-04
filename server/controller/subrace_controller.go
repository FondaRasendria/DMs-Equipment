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

type SubraceController struct {
	SubraceService service.SubraceService
}

func NewSubraceController(subraceService service.SubraceService) *SubraceController {
	return &SubraceController{SubraceService: subraceService}
}

func (controller *SubraceController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	subraceCreateRequest := request.SubraceCreateRequest{}
	helper.ReadRequestBody(requests, &subraceCreateRequest)

	result := controller.SubraceService.Create(requests.Context(), subraceCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SubraceController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	subraceUpdateRequest := request.SubraceUpdateRequest{}
	helper.ReadRequestBody(requests, &subraceUpdateRequest)

	subraceId := params.ByName("subraceId")
	id, err := strconv.Atoi(subraceId)
	helper.PanicIfError(err)
	subraceUpdateRequest.Id = id
	result := controller.SubraceService.Update(requests.Context(), subraceUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SubraceController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	subraceId := params.ByName("subraceId")
	id, err := strconv.Atoi(subraceId)
	helper.PanicIfError(err)

	result := controller.SubraceService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SubraceController) SelectById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	subraceId := params.ByName("subraceId")
	id, err := strconv.Atoi(subraceId)
	helper.PanicIfError(err)

	result := controller.SubraceService.SelectById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SubraceController) SelectAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.SubraceService.SelectAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SubraceController) SelectByParentId(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	parentId := params.ByName("parentId")
	id, err := strconv.Atoi(parentId)
	helper.PanicIfError(err)

	result := controller.SubraceService.SelectByParentId(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
