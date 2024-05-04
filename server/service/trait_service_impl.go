package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
)

type TraitServiceImpl struct {
	TraitRepository repository.TraitRepository
}

// Create implements TraitService.
func (s *TraitServiceImpl) Create(ctx context.Context, request request.TraitCreateRequest) response.TraitResponse {
	trait := model.Trait{
		ParentId:    request.ParentId,
		Name:        request.Name,
		Description: request.Description,
	}
	s.TraitRepository.Create(ctx, trait)
	return response.TraitResponse(trait)
}

// Delete implements TraitService.
func (s *TraitServiceImpl) Delete(ctx context.Context, traitId int) response.TraitResponse {
	trait, err := s.TraitRepository.SelectById(ctx, traitId)
	helper.PanicIfError(err)
	s.TraitRepository.Delete(ctx, trait.Id)
	return response.TraitResponse(trait)
}

// SelectAll implements TraitService.
func (s *TraitServiceImpl) SelectAll(ctx context.Context) []response.TraitResponse {
	traits := s.TraitRepository.SelectAll(ctx)

	var traitsResp []response.TraitResponse

	for _, value := range traits {
		trait := response.TraitResponse{Id: value.Id, ParentId: value.ParentId, Name: value.Name, Description: value.Description}
		traitsResp = append(traitsResp, trait)
	}
	return traitsResp
}

// SelectById implements TraitService.
func (s *TraitServiceImpl) SelectById(ctx context.Context, traitId int) response.TraitResponse {
	trait, err := s.TraitRepository.SelectById(ctx, traitId)
	helper.PanicIfError(err)
	return response.TraitResponse(trait)
}

// Update implements TraitService.
func (s *TraitServiceImpl) Update(ctx context.Context, request request.TraitUpdateRequest) response.TraitResponse {
	trait, err := s.TraitRepository.SelectById(ctx, request.Id)
	helper.PanicIfError(err)

	trait.ParentId = request.ParentId
	trait.Name = request.Name
	trait.Description = request.Description
	s.TraitRepository.Update(ctx, trait)
	return response.TraitResponse(trait)
}

// SelectByParentId implements TraitService.
func (s *TraitServiceImpl) SelectByParentId(ctx context.Context, paarentId int) []response.TraitResponse {
	traits := s.TraitRepository.SelectByParentId(ctx, paarentId)

	var traitsResp []response.TraitResponse

	for _, value := range traits {
		trait := response.TraitResponse{Id: value.Id, ParentId: value.ParentId, Name: value.Name, Description: value.Description}
		traitsResp = append(traitsResp, trait)
	}
	return traitsResp
}

func NewTraitServiceImpl(traitRepository repository.TraitRepository) TraitService {
	return &TraitServiceImpl{TraitRepository: traitRepository}
}
