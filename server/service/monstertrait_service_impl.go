package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
)

type MonsterTraitServiceImpl struct {
	MonsterTraitRepository repository.MonsterTraitRepository
}

// Create implements MonsterTraitService.
func (s *MonsterTraitServiceImpl) Create(ctx context.Context, request request.MonsterTraitCreateRequest) response.MonsterTraitResponse {
	mt := model.MonsterTrait{
		ParentId:    request.ParentId,
		Name:        request.Name,
		Description: request.Description,
	}
	s.MonsterTraitRepository.Create(ctx, mt)
	return response.MonsterTraitResponse(mt)
}

// Delete implements MonsterTraitService.
func (s *MonsterTraitServiceImpl) Delete(ctx context.Context, mtId int) response.MonsterTraitResponse {
	mt, err := s.MonsterTraitRepository.SelectById(ctx, mtId)
	helper.PanicIfError(err)
	s.MonsterTraitRepository.Delete(ctx, mt.Id)
	return response.MonsterTraitResponse(mt)
}

// SelectAll implements MonsterTraitService.
func (s *MonsterTraitServiceImpl) SelectAll(ctx context.Context) []response.MonsterTraitResponse {
	mts := s.MonsterTraitRepository.SelectAll(ctx)

	var mtsResp []response.MonsterTraitResponse

	for _, value := range mts {
		mt := response.MonsterTraitResponse{Id: value.Id, ParentId: value.ParentId, Name: value.Name, Description: value.Description}
		mtsResp = append(mtsResp, mt)
	}

	return mtsResp
}

// SelectById implements MonsterTraitService.
func (s *MonsterTraitServiceImpl) SelectById(ctx context.Context, mtId int) response.MonsterTraitResponse {
	mt, err := s.MonsterTraitRepository.SelectById(ctx, mtId)
	helper.PanicIfError(err)
	return response.MonsterTraitResponse(mt)
}

// SelectByParentId implements MonsterTraitService.
func (s *MonsterTraitServiceImpl) SelectByParentId(ctx context.Context, parentId int) []response.MonsterTraitResponse {
	mts := s.MonsterTraitRepository.SelectByParentId(ctx, parentId)

	var mtsResp []response.MonsterTraitResponse

	for _, value := range mts {
		mt := response.MonsterTraitResponse{Id: value.Id, ParentId: value.ParentId, Name: value.Name, Description: value.Description}
		mtsResp = append(mtsResp, mt)
	}

	return mtsResp
}

// Update implements MonsterTraitService.
func (s *MonsterTraitServiceImpl) Update(ctx context.Context, request request.MonsterTraitUpdateRequest) response.MonsterTraitResponse {
	mt, err := s.MonsterTraitRepository.SelectById(ctx, request.Id)
	helper.PanicIfError(err)

	mt.ParentId = request.ParentId
	mt.Name = request.Name
	mt.Description = request.Description
	s.MonsterTraitRepository.Update(ctx, mt)
	return response.MonsterTraitResponse(mt)
}

func NewMonsterTraitServiceImpl(mtRepository repository.MonsterTraitRepository) MonsterTraitService {
	return &MonsterTraitServiceImpl{MonsterTraitRepository: mtRepository}
}
