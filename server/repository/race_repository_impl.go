package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud/helper"
	"golang-crud/model"
)

type RaceRepositoryImpl struct {
	Db *sql.DB
}

// Create implements RaceRepository.
func (r *RaceRepositoryImpl) Create(ctx context.Context, race model.Race) model.Race {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tx2, err2 := r.Db.Begin()
	helper.PanicIfError(err2)
	defer helper.CommitOrRollback(tx2)

	SQLid := "select max(id) from race"
	result, errQuery := tx.QueryContext(ctx, SQLid)
	helper.PanicIfError(errQuery)
	defer result.Close()

	maxId := 0

	if result.Next() {
		err := result.Scan(&maxId)
		maxId = maxId + 1
		if err != nil {
			maxId = 1
		}
	}

	SQL := "insert into race(id, name) values ($1, $2)"
	_, errExec := tx2.ExecContext(ctx, SQL, maxId, race.Name)
	helper.PanicIfError(errExec)

	return model.Race{Id: maxId, Name: race.Name}
}

// Delete implements RaceRepository.
func (r *RaceRepositoryImpl) Delete(ctx context.Context, raceId int) model.Race {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	race, errSelect := r.SelectById(ctx, raceId)
	helper.PanicIfError(errSelect)

	SQL := "delete from race where id=$1"
	_, errExec := tx.ExecContext(ctx, SQL, raceId)
	helper.PanicIfError(errExec)
	return race
}

// SelectAll implements RaceRepository.
func (r *RaceRepositoryImpl) SelectAll(ctx context.Context) []model.Race {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from race"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var races []model.Race

	for result.Next() {
		race := model.Race{}
		err := result.Scan(&race.Id, &race.Name)
		helper.PanicIfError(err)

		races = append(races, race)
	}
	return races
}

// SelectById implements RaceRepository.
func (r *RaceRepositoryImpl) SelectById(ctx context.Context, raceId int) (model.Race, error) {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from race where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, raceId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	race := model.Race{}

	if result.Next() {
		err := result.Scan(&race.Id, &race.Name)
		helper.PanicIfError(err)
		return race, nil
	} else {
		return race, errors.New("race with that id not found")
	}
}

// Update implements RaceRepository.
func (r *RaceRepositoryImpl) Update(ctx context.Context, race model.Race) model.Race {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update race set name=$1 where id=$2"
	_, errExec := tx.ExecContext(ctx, SQL, race.Name, race.Id)
	helper.PanicIfError(errExec)

	return race
}

// SelectByName implements RaceRepository.
func (r *RaceRepositoryImpl) SelectByName(ctx context.Context, name string) (model.Race, error) {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from race where name=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(errQuery)
	defer result.Close()

	race := model.Race{}

	if result.Next() {
		err := result.Scan(&race.Id, &race.Name)
		helper.PanicIfError(err)
		return race, nil
	} else {
		return race, errors.New("race with that name not found")
	}
}

func NewRaceRepository(Db *sql.DB) RaceRepository {
	return &RaceRepositoryImpl{Db: Db}
}
