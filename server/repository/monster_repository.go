package repository

import (
	"context"
	"golang-crud/model"
)

type MonsterRepository interface {
	Create(ctx context.Context, monster model.Monster) model.Monster
	Update(ctx context.Context, monster model.Monster) model.Monster
	Delete(ctx context.Context, monsterId int) model.Monster
	SelectById(ctx context.Context, monsterId int) (model.Monster, error)
	SelectAll(ctx context.Context) []model.Monster

	SelectByName(ctx context.Context, name string) (model.Monster, error)
	SelectBySource(ctx context.Context, source string) []model.Monster
}
