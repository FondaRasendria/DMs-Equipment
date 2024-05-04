package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
)

type MonsterTraitService interface {
	Create(ctx context.Context, request request.MonsterTraitCreateRequest) response.MonsterTraitResponse
	Update(ctx context.Context, request request.MonsterTraitUpdateRequest) response.MonsterTraitResponse
	Delete(ctx context.Context, mtId int) response.MonsterTraitResponse
	SelectById(ctx context.Context, mtId int) response.MonsterTraitResponse
	SelectAll(ctx context.Context) []response.MonsterTraitResponse

	SelectByParentId(ctx context.Context, parentId int) []response.MonsterTraitResponse
}
