package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud/helper"
	"golang-crud/model"
)

type MonsterTraitRepositoryImpl struct {
	Db *sql.DB
}

// Create implements MonsterTraitRepository.
func (r *MonsterTraitRepositoryImpl) Create(ctx context.Context, mt model.MonsterTrait) model.MonsterTrait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tx2, err2 := r.Db.Begin()
	helper.PanicIfError(err2)
	defer helper.CommitOrRollback(tx2)

	SQLid := "select max(id) from monstertrait"
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

	SQL := "insert into monstertrait(id, parentid, name, description) values ($1, $2, $3, $4)"
	_, errExec := tx2.ExecContext(ctx, SQL, maxId, mt.ParentId, mt.Name, mt.Description)
	helper.PanicIfError(errExec)
	return model.MonsterTrait{Id: maxId, ParentId: mt.ParentId, Name: mt.Name, Description: mt.Description}
}

// Delete implements MonsterTraitRepository.
func (r *MonsterTraitRepositoryImpl) Delete(ctx context.Context, mtId int) model.MonsterTrait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mt, errSelect := r.SelectById(ctx, mtId)
	helper.PanicIfError(errSelect)

	SQL := "delete from monstertrait where id=$1"
	_, errExec := tx.ExecContext(ctx, SQL, mtId)
	helper.PanicIfError(errExec)
	return mt
}

// SelectAll implements MonsterTraitRepository.
func (r *MonsterTraitRepositoryImpl) SelectAll(ctx context.Context) []model.MonsterTrait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from monstertrait"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var mts []model.MonsterTrait

	for result.Next() {
		mt := model.MonsterTrait{}
		err := result.Scan(&mt.Id, &mt.ParentId, &mt.Name, &mt.Description)
		helper.PanicIfError(err)

		mts = append(mts, mt)
	}
	return mts
}

// SelectById implements MonsterTraitRepository.
func (r *MonsterTraitRepositoryImpl) SelectById(ctx context.Context, mtId int) (model.MonsterTrait, error) {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from monstertrait where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, mtId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	mt := model.MonsterTrait{}

	if result.Next() {
		err := result.Scan(&mt.Id, &mt.ParentId, &mt.Name, &mt.Description)
		helper.PanicIfError(err)
		return mt, nil
	} else {
		return mt, errors.New("race with that id not found")
	}
}

// SelectByParentId implements MonsterTraitRepository.
func (r *MonsterTraitRepositoryImpl) SelectByParentId(ctx context.Context, parentId int) []model.MonsterTrait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from monstertrait where parentid=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, parentId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var mts []model.MonsterTrait

	for result.Next() {
		mt := model.MonsterTrait{}
		err := result.Scan(&mt.Id, &mt.ParentId, &mt.Name, &mt.Description)
		helper.PanicIfError(err)

		mts = append(mts, mt)
	}
	return mts
}

// Update implements MonsterTraitRepository.
func (r *MonsterTraitRepositoryImpl) Update(ctx context.Context, mt model.MonsterTrait) model.MonsterTrait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update monstertrait set parentid=$1,name=$2,description=$3 where id=$4"
	_, errExec := tx.ExecContext(ctx, SQL, mt.ParentId, mt.Name, mt.Description, mt.Id)
	helper.PanicIfError(errExec)
	return mt
}

func NewMonsterTraitRepositoryImpl(Db *sql.DB) MonsterTraitRepository {
	return &MonsterTraitRepositoryImpl{Db: Db}
}
