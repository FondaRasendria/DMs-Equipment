package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
)

type SubraceServiceImpl struct {
	SubraceRepository repository.SubraceRepository
}

// Create implements SubraceService.
func (s *SubraceServiceImpl) Create(ctx context.Context, request request.SubraceCreateRequest) response.SubraceResponse {
	subrace := model.Subrace{
		ParentId:    request.ParentId,
		Description: request.Description,
		Source:      request.Source,
		Ability:     request.Ability,
		Age:         request.Age,
		Alignment:   request.Alignment,
		Size:        request.Size,
		Speed:       request.Speed,
		Language:    request.Language,
	}
	result := s.SubraceRepository.Create(ctx, subrace)
	return response.SubraceResponse(result)
}

// Delete implements SubraceService.
func (s *SubraceServiceImpl) Delete(ctx context.Context, subraceId int) response.SubraceResponse {
	subrace, err := s.SubraceRepository.SelectById(ctx, subraceId)
	helper.PanicIfError(err)
	result := s.SubraceRepository.Delete(ctx, subrace.Id)
	return response.SubraceResponse(result)
}

// SelectAll implements SubraceService.
func (s *SubraceServiceImpl) SelectAll(ctx context.Context) []response.SubraceResponse {
	subraces := s.SubraceRepository.SelectAll(ctx)

	var subracesResp []response.SubraceResponse

	for _, value := range subraces {
		subrace := response.SubraceResponse{Id: value.Id, ParentId: value.ParentId, Description: value.Description, Source: value.Source, Ability: value.Ability, Age: value.Age, Alignment: value.Alignment, Size: value.Size, Speed: value.Speed, Language: value.Language}
		subracesResp = append(subracesResp, subrace)
	}

	return subracesResp
}

// SelectById implements SubraceService.
func (s *SubraceServiceImpl) SelectById(ctx context.Context, subraceId int) response.SubraceResponse {
	subrace, err := s.SubraceRepository.SelectById(ctx, subraceId)
	helper.PanicIfError(err)
	return response.SubraceResponse(subrace)
}

// Update implements SubraceService.
func (s *SubraceServiceImpl) Update(ctx context.Context, request request.SubraceUpdateRequest) response.SubraceResponse {
	subrace, err := s.SubraceRepository.SelectById(ctx, request.Id)
	helper.PanicIfError(err)

	subrace.ParentId = request.ParentId
	subrace.Description = request.Description
	subrace.Source = request.Source
	subrace.Ability = request.Ability
	subrace.Age = request.Age
	subrace.Alignment = request.Alignment
	subrace.Size = request.Size
	subrace.Speed = request.Speed
	subrace.Language = request.Language
	result := s.SubraceRepository.Update(ctx, subrace)
	return response.SubraceResponse(result)
}

// SelectByParentId implements SubraceService.
func (s *SubraceServiceImpl) SelectByParentId(ctx context.Context, parentId int) []response.SubraceResponse {
	subraces := s.SubraceRepository.SelectByParentId(ctx, parentId)

	var subracesResp []response.SubraceResponse

	for _, value := range subraces {
		subrace := response.SubraceResponse{Id: value.Id, ParentId: value.ParentId, Description: value.Description, Source: value.Source, Ability: value.Ability, Age: value.Age, Alignment: value.Alignment, Size: value.Size, Speed: value.Speed, Language: value.Language}
		subracesResp = append(subracesResp, subrace)
	}

	return subracesResp
}

func NewSubraceServiceImpl(subraceRepository repository.SubraceRepository) SubraceService {
	return &SubraceServiceImpl{SubraceRepository: subraceRepository}
}
