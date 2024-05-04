package repository

import (
	"context"
	"golang-crud/model"
)

type RaceRepository interface {
	Create(ctx context.Context, race model.Race) model.Race
	Update(ctx context.Context, race model.Race) model.Race
	Delete(ctx context.Context, raceId int) model.Race
	SelectById(ctx context.Context, raceId int) (model.Race, error)
	SelectAll(ctx context.Context) []model.Race

	SelectByName(ctx context.Context, name string) (model.Race, error)
}
