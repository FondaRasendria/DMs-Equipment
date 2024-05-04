package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
)

type ActionService interface {
	Create(ctx context.Context, request request.ActionCreateRequest) response.ActionResponse
	Update(ctx context.Context, request request.ActionUpdateRequest) response.ActionResponse
	Delete(ctx context.Context, actionId int) response.ActionResponse
	SelectById(ctx context.Context, actionId int) response.ActionResponse
	SelectAll(ctx context.Context) []response.ActionResponse

	SelectByParentId(ctx context.Context, parentId int) []response.ActionResponse
}
