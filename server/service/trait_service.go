package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
)

type TraitService interface {
	Create(ctx context.Context, request request.TraitCreateRequest) response.TraitResponse
	Update(ctx context.Context, request request.TraitUpdateRequest) response.TraitResponse
	Delete(ctx context.Context, traitId int) response.TraitResponse
	SelectById(ctx context.Context, traitId int) response.TraitResponse
	SelectAll(ctx context.Context) []response.TraitResponse

	SelectByParentId(ctx context.Context, paarentId int) []response.TraitResponse
}
