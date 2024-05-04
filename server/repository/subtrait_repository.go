package repository

import (
	"context"
	"golang-crud/model"
)

type SubtraitRepository interface {
	Create(ctx context.Context, subtrait model.Subtrait) model.Subtrait
	Update(ctx context.Context, subtrait model.Subtrait) model.Subtrait
	Delete(ctx context.Context, subtraitId int) model.Subtrait
	SelectById(ctx context.Context, subtraitId int) (model.Subtrait, error)
	SelectAll(ctx context.Context) []model.Subtrait

	SelectByParentId(ctx context.Context, parentId int) []model.Subtrait
}
