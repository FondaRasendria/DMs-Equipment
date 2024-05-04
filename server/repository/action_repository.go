package repository

import (
	"context"
	"golang-crud/model"
)

type ActionRepository interface {
	Create(ctx context.Context, action model.Action) model.Action
	Update(ctx context.Context, action model.Action) model.Action
	Delete(ctx context.Context, actionId int) model.Action
	SelectById(ctx context.Context, actionId int) (model.Action, error)
	SelectAll(ctx context.Context) []model.Action

	SelectByParentId(ctx context.Context, parentId int) []model.Action
}
