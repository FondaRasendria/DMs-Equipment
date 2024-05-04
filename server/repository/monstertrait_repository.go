package repository

import (
	"context"
	"golang-crud/model"
)

type MonsterTraitRepository interface {
	Create(ctx context.Context, mt model.MonsterTrait) model.MonsterTrait
	Update(ctx context.Context, mt model.MonsterTrait) model.MonsterTrait
	Delete(ctx context.Context, mtId int) model.MonsterTrait
	SelectById(ctx context.Context, mtId int) (model.MonsterTrait, error)
	SelectAll(ctx context.Context) []model.MonsterTrait

	SelectByParentId(ctx context.Context, parentId int) []model.MonsterTrait
}
