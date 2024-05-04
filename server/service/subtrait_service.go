package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
)

type SubtraitService interface {
	Create(ctx context.Context, request request.SubtraitCreateRequest) response.SubtraitResponse
	Update(ctx context.Context, request request.SubtraitUpdateRequest) response.SubtraitResponse
	Delete(ctx context.Context, subtraitId int) response.SubtraitResponse
	SelectById(ctx context.Context, subtraitId int) response.SubtraitResponse
	SelectAll(ctx context.Context) []response.SubtraitResponse

	SelectByParentId(ctx context.Context, parentId int) []response.SubtraitResponse
}
