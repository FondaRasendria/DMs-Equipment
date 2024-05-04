package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud/helper"
	"golang-crud/model"
)

type ActionRepositoryImpl struct {
	Db *sql.DB
}

// Create implements ActionRepository.
func (r *ActionRepositoryImpl) Create(ctx context.Context, action model.Action) model.Action {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tx2, err2 := r.Db.Begin()
	helper.PanicIfError(err2)
	defer helper.CommitOrRollback(tx2)

	SQLid := "select max(id) from action"
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

	SQL := "insert into action(id, parentid, name, description) values ($1, $2, $3, $4)"
	_, errExec := tx2.ExecContext(ctx, SQL, maxId, action.ParentId, action.Name, action.Description)
	helper.PanicIfError(errExec)
	return model.Action{Id: maxId, ParentId: action.ParentId, Name: action.Name, Description: action.Description}
}

// Delete implements ActionRepository.
func (r *ActionRepositoryImpl) Delete(ctx context.Context, actionId int) model.Action {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	action, errSelect := r.SelectById(ctx, actionId)
	helper.PanicIfError(errSelect)

	SQL := "delete from action where id=$1"
	_, errExec := tx.ExecContext(ctx, SQL, actionId)
	helper.PanicIfError(errExec)
	return action
}

// SelectAll implements ActionRepository.
func (r *ActionRepositoryImpl) SelectAll(ctx context.Context) []model.Action {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from action"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var actions []model.Action

	for result.Next() {
		action := model.Action{}
		err := result.Scan(&action.Id, &action.ParentId, &action.Name, &action.Description)
		helper.PanicIfError(err)

		actions = append(actions, action)
	}
	return actions
}

// SelectById implements ActionRepository.
func (r *ActionRepositoryImpl) SelectById(ctx context.Context, actionId int) (model.Action, error) {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from action where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, actionId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	action := model.Action{}

	if result.Next() {
		err := result.Scan(&action.Id, &action.ParentId, &action.Name, &action.Description)
		helper.PanicIfError(err)
		return action, nil
	} else {
		return action, errors.New("race with that id not found")
	}
}

// SelectByParentId implements ActionRepository.
func (r *ActionRepositoryImpl) SelectByParentId(ctx context.Context, parentId int) []model.Action {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from action where parentid=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, parentId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var actions []model.Action

	for result.Next() {
		action := model.Action{}
		err := result.Scan(&action.Id, &action.ParentId, &action.Name, &action.Description)
		helper.PanicIfError(err)

		actions = append(actions, action)
	}
	return actions
}

// Update implements ActionRepository.
func (r *ActionRepositoryImpl) Update(ctx context.Context, action model.Action) model.Action {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update action set parentid=$1,name=$2,description=$3 where id=$4"
	_, errExec := tx.ExecContext(ctx, SQL, action.ParentId, action.Name, action.Description, action.Id)
	helper.PanicIfError(errExec)
	return action
}

func NewActionRepository(Db *sql.DB) ActionRepository {
	return &ActionRepositoryImpl{Db: Db}
}
