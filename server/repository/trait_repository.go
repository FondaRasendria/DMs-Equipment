package repository

import (
	"context"
	"golang-crud/model"
)

type TraitRepository interface {
	Create(ctx context.Context, trait model.Trait) model.Trait
	Update(ctx context.Context, trait model.Trait) model.Trait
	Delete(ctx context.Context, traitId int) model.Trait
	SelectById(ctx context.Context, traitId int) (model.Trait, error)
	SelectAll(ctx context.Context) []model.Trait

	SelectByParentId(ctx context.Context, parentId int) []model.Trait
}
