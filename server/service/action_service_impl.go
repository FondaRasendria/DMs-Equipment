package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
)

type ActionServiceImpl struct {
	ActionRepository repository.ActionRepository
}

// Create implements ActionService.
func (s *ActionServiceImpl) Create(ctx context.Context, request request.ActionCreateRequest) response.ActionResponse {
	action := model.Action{
		ParentId:    request.ParentId,
		Name:        request.Name,
		Description: request.Description,
	}
	s.ActionRepository.Create(ctx, action)
	return response.ActionResponse(action)
}

// Delete implements ActionService.
func (s *ActionServiceImpl) Delete(ctx context.Context, actionId int) response.ActionResponse {
	action, err := s.ActionRepository.SelectById(ctx, actionId)
	helper.PanicIfError(err)
	s.ActionRepository.Delete(ctx, action.Id)
	return response.ActionResponse(action)
}

// SelectAll implements ActionService.
func (s *ActionServiceImpl) SelectAll(ctx context.Context) []response.ActionResponse {
	actions := s.ActionRepository.SelectAll(ctx)

	var actionsResp []response.ActionResponse

	for _, value := range actions {
		action := response.ActionResponse{Id: value.Id, ParentId: value.ParentId, Name: value.Name, Description: value.Description}
		actionsResp = append(actionsResp, action)
	}

	return actionsResp
}

// SelectById implements ActionService.
func (s *ActionServiceImpl) SelectById(ctx context.Context, actionId int) response.ActionResponse {
	action, err := s.ActionRepository.SelectById(ctx, actionId)
	helper.PanicIfError(err)
	return response.ActionResponse(action)
}

// SelectByParentId implements ActionService.
func (s *ActionServiceImpl) SelectByParentId(ctx context.Context, parentId int) []response.ActionResponse {
	actions := s.ActionRepository.SelectByParentId(ctx, parentId)

	var actionsResp []response.ActionResponse

	for _, value := range actions {
		action := response.ActionResponse{Id: value.Id, ParentId: value.ParentId, Name: value.Name, Description: value.Description}
		actionsResp = append(actionsResp, action)
	}

	return actionsResp
}

// Update implements ActionService.
func (s *ActionServiceImpl) Update(ctx context.Context, request request.ActionUpdateRequest) response.ActionResponse {
	action, err := s.ActionRepository.SelectById(ctx, request.Id)
	helper.PanicIfError(err)

	action.ParentId = request.ParentId
	action.Name = request.Name
	action.Description = request.Description
	s.ActionRepository.Update(ctx, action)
	return response.ActionResponse(action)
}

func NewActionServiceImpl(actionRepository repository.ActionRepository) ActionService {
	return &ActionServiceImpl{ActionRepository: actionRepository}
}
