package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
)

type MonsterService interface {
	Create(ctx context.Context, request request.MonsterCreateRequest) response.MonsterResponse
	Update(ctx context.Context, request request.MonsterUpdateRequest) response.MonsterResponse
	Delete(ctx context.Context, monsterId int) response.MonsterResponse
	SelectById(ctx context.Context, monsterId int) response.MonsterResponse
	SelectAll(ctx context.Context) []response.MonsterResponse

	SelectByName(ctx context.Context, name string) response.MonsterResponse
	SelectBySource(ctx context.Context, source string) []response.MonsterResponse
}
