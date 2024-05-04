package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud/helper"
	"golang-crud/model"
)

type SubtraitRepositoryImpl struct {
	Db *sql.DB
}

// Create implements SubtraitRepository.
func (r *SubtraitRepositoryImpl) Create(ctx context.Context, subtrait model.Subtrait) model.Subtrait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tx2, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx2)

	SQLid := "select max(id) from subtrait"
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

	SQL := "insert into subtrait(id, parentid, level, description) values ($1, $2, $3, $4)"
	_, errExec := tx2.ExecContext(ctx, SQL, maxId, subtrait.ParentId, subtrait.Level, subtrait.Description)
	helper.PanicIfError(errExec)
	return model.Subtrait{Id: maxId, ParentId: subtrait.ParentId, Level: subtrait.Level, Description: subtrait.Description}
}

// Delete implements SubtraitRepository.
func (r *SubtraitRepositoryImpl) Delete(ctx context.Context, subtraitId int) model.Subtrait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	subtrait, errSelect := r.SelectById(ctx, subtraitId)
	helper.PanicIfError(errSelect)

	SQL := "delete from subtrait where id=$1"
	_, errExec := tx.ExecContext(ctx, SQL, subtraitId)
	helper.PanicIfError(errExec)
	return subtrait
}

// SelectAll implements SubtraitRepository.
func (r *SubtraitRepositoryImpl) SelectAll(ctx context.Context) []model.Subtrait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from subtrait"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var subtraits []model.Subtrait

	for result.Next() {
		subtrait := model.Subtrait{}
		err := result.Scan(&subtrait.Id, &subtrait.ParentId, &subtrait.Level, &subtrait.Description)
		helper.PanicIfError(err)

		subtraits = append(subtraits, subtrait)
	}
	return subtraits
}

// SelectById implements SubtraitRepository.
func (r *SubtraitRepositoryImpl) SelectById(ctx context.Context, subtraitId int) (model.Subtrait, error) {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from subtrait where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, subtraitId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	subtrait := model.Subtrait{}

	if result.Next() {
		err := result.Scan(&subtrait.Id, &subtrait.ParentId, &subtrait.Level, &subtrait.Description)
		helper.PanicIfError(err)
		return subtrait, nil
	} else {
		return subtrait, errors.New("subtrait with that id not found")
	}
}

// Update implements SubtraitRepository.
func (r *SubtraitRepositoryImpl) Update(ctx context.Context, subtrait model.Subtrait) model.Subtrait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update subtrait set parentid=$1,level=$2,description=$3 where id=$4"
	_, errExec := tx.ExecContext(ctx, SQL, subtrait.ParentId, subtrait.Level, subtrait.Description, subtrait.Id)
	helper.PanicIfError(errExec)
	return subtrait
}

// SelectByParentId implements SubtraitRepository.
func (r *SubtraitRepositoryImpl) SelectByParentId(ctx context.Context, parentId int) []model.Subtrait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from subtrait where parentId = $1"
	result, errQuery := tx.QueryContext(ctx, SQL, parentId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var subtraits []model.Subtrait

	for result.Next() {
		subtrait := model.Subtrait{}
		err := result.Scan(&subtrait.Id, &subtrait.ParentId, &subtrait.Level, &subtrait.Description)
		helper.PanicIfError(err)

		subtraits = append(subtraits, subtrait)
	}
	return subtraits
}

func NewSubtraitRepository(Db *sql.DB) SubtraitRepository {
	return &SubtraitRepositoryImpl{Db: Db}
}
