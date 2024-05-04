package repository

import (
	"context"
	"golang-crud/model"
)

type SubraceRepository interface {
	Create(ctx context.Context, subrace model.Subrace) model.Subrace
	Update(ctx context.Context, subrace model.Subrace) model.Subrace
	Delete(ctx context.Context, subraceId int) model.Subrace
	SelectById(ctx context.Context, subraceId int) (model.Subrace, error)
	SelectAll(ctx context.Context) []model.Subrace

	SelectByParentId(ctx context.Context, parentId int) []model.Subrace
}
