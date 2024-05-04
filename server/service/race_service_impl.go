package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
)

type RaceServiceImpl struct {
	RaceRepository repository.RaceRepository
}

// Create implements RaceService.
func (s *RaceServiceImpl) Create(ctx context.Context, request request.RaceCreateRequest) response.RaceResponse {
	race := model.Race{
		Name: request.Name,
	}
	result := s.RaceRepository.Create(ctx, race)
	return response.RaceResponse(result)
}

// Delete implements RaceService.
func (s *RaceServiceImpl) Delete(ctx context.Context, raceId int) response.RaceResponse {
	race, err := s.RaceRepository.SelectById(ctx, raceId)
	helper.PanicIfError(err)
	s.RaceRepository.Delete(ctx, race.Id)
	return response.RaceResponse(race)
}

// SelectAll implements RaceService.
func (s *RaceServiceImpl) SelectAll(ctx context.Context) []response.RaceResponse {
	races := s.RaceRepository.SelectAll(ctx)

	var racesResp []response.RaceResponse

	for _, value := range races {
		race := response.RaceResponse{Id: value.Id, Name: value.Name}
		racesResp = append(racesResp, race)
	}

	return racesResp
}

// SelectById implements RaceService.
func (s *RaceServiceImpl) SelectById(ctx context.Context, raceId int) response.RaceResponse {
	race, err := s.RaceRepository.SelectById(ctx, raceId)
	helper.PanicIfError(err)
	return response.RaceResponse(race)
}

// Update implements RaceService.
func (s *RaceServiceImpl) Update(ctx context.Context, request request.RaceUpdateRequest) response.RaceResponse {
	race, err := s.RaceRepository.SelectById(ctx, request.Id)
	helper.PanicIfError(err)

	race.Name = request.Name
	s.RaceRepository.Update(ctx, race)
	return response.RaceResponse(race)
}

// SelectByName implements RaceService.
func (s *RaceServiceImpl) SelectByName(ctx context.Context, name string) response.RaceResponse {
	race, err := s.RaceRepository.SelectByName(ctx, name)
	helper.PanicIfError(err)
	return response.RaceResponse(race)
}

func NewRaceServiceImpl(raceRepository repository.RaceRepository) RaceService {
	return &RaceServiceImpl{RaceRepository: raceRepository}
}
