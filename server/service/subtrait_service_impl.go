package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
)

type SubtraitServiceImpl struct {
	SubtraitRepository repository.SubtraitRepository
}

// Create implements SubtraitService.
func (s *SubtraitServiceImpl) Create(ctx context.Context, request request.SubtraitCreateRequest) response.SubtraitResponse {
	subtrait := model.Subtrait{
		ParentId:    request.ParentId,
		Level:       request.Level,
		Description: request.Description,
	}
	s.SubtraitRepository.Create(ctx, subtrait)
	return response.SubtraitResponse(subtrait)
}

// Delete implements SubtraitService.
func (s *SubtraitServiceImpl) Delete(ctx context.Context, subtraitId int) response.SubtraitResponse {
	subtrait, err := s.SubtraitRepository.SelectById(ctx, subtraitId)
	helper.PanicIfError(err)
	s.SubtraitRepository.Delete(ctx, subtrait.Id)
	return response.SubtraitResponse(subtrait)
}

// SelectAll implements SubtraitService.
func (s *SubtraitServiceImpl) SelectAll(ctx context.Context) []response.SubtraitResponse {
	subtraits := s.SubtraitRepository.SelectAll(ctx)

	var subtraitsResp []response.SubtraitResponse

	for _, value := range subtraits {
		subtrait := response.SubtraitResponse{Id: value.Id, ParentId: value.ParentId, Level: value.Level, Description: value.Description}
		subtraitsResp = append(subtraitsResp, subtrait)
	}

	return subtraitsResp
}

// SelectById implements SubtraitService.
func (s *SubtraitServiceImpl) SelectById(ctx context.Context, subtraitId int) response.SubtraitResponse {
	subtrait, err := s.SubtraitRepository.SelectById(ctx, subtraitId)
	helper.PanicIfError(err)
	return response.SubtraitResponse(subtrait)
}

// Update implements SubtraitService.
func (s *SubtraitServiceImpl) Update(ctx context.Context, request request.SubtraitUpdateRequest) response.SubtraitResponse {
	subtrait, err := s.SubtraitRepository.SelectById(ctx, request.Id)
	helper.PanicIfError(err)

	subtrait.ParentId = request.ParentId
	subtrait.Level = request.Level
	subtrait.Description = request.Description
	s.SubtraitRepository.Update(ctx, subtrait)
	return response.SubtraitResponse(subtrait)
}

// SelectByParentId implements SubtraitService.
func (s *SubtraitServiceImpl) SelectByParentId(ctx context.Context, parentId int) []response.SubtraitResponse {
	subtraits := s.SubtraitRepository.SelectByParentId(ctx, parentId)

	var subtraitsResp []response.SubtraitResponse

	for _, value := range subtraits {
		subtrait := response.SubtraitResponse{Id: value.Id, ParentId: value.ParentId, Level: value.Level, Description: value.Description}
		subtraitsResp = append(subtraitsResp, subtrait)
	}

	return subtraitsResp
}

func NewSubtraitServiceImpl(subtraitRepository repository.SubtraitRepository) SubtraitService {
	return &SubtraitServiceImpl{SubtraitRepository: subtraitRepository}
}
