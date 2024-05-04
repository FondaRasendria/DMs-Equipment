package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
)

type RaceService interface {
	Create(ctx context.Context, request request.RaceCreateRequest) response.RaceResponse
	Update(ctx context.Context, request request.RaceUpdateRequest) response.RaceResponse
	Delete(ctx context.Context, raceId int) response.RaceResponse
	SelectById(ctx context.Context, raceId int) response.RaceResponse
	SelectAll(ctx context.Context) []response.RaceResponse

	SelectByName(ctx context.Context, name string) response.RaceResponse
}
