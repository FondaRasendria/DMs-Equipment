package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
)

type SubraceService interface {
	Create(ctx context.Context, request request.SubraceCreateRequest) response.SubraceResponse
	Update(ctx context.Context, request request.SubraceUpdateRequest) response.SubraceResponse
	Delete(ctx context.Context, subraceId int) response.SubraceResponse
	SelectById(ctx context.Context, subraceId int) response.SubraceResponse
	SelectAll(ctx context.Context) []response.SubraceResponse

	SelectByParentId(ctx context.Context, parentId int) []response.SubraceResponse
}
