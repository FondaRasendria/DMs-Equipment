package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud/helper"
	"golang-crud/model"
)

type TraitRepositoryImpl struct {
	Db *sql.DB
}

// Create implements TraitRepository.
func (r *TraitRepositoryImpl) Create(ctx context.Context, trait model.Trait) model.Trait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tx2, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx2)

	SQLid := "select max(id) from trait"
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

	SQL := "insert into trait(id, parentid, name, description) values ($1, $2, $3, $4)"
	_, errExec := tx2.ExecContext(ctx, SQL, maxId, trait.ParentId, trait.Name, trait.Description)
	helper.PanicIfError(errExec)
	return model.Trait{Id: maxId, ParentId: trait.ParentId, Name: trait.Name, Description: trait.Description}
}

// Delete implements TraitRepository.
func (r *TraitRepositoryImpl) Delete(ctx context.Context, traitId int) model.Trait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	trait, errSelect := r.SelectById(ctx, traitId)
	helper.PanicIfError(errSelect)

	SQL := "delete from trait where id=$1"
	_, errExec := tx.ExecContext(ctx, SQL, traitId)
	helper.PanicIfError(errExec)
	return trait
}

// SelectAll implements TraitRepository.
func (r *TraitRepositoryImpl) SelectAll(ctx context.Context) []model.Trait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from trait"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var traits []model.Trait

	for result.Next() {
		trait := model.Trait{}
		err := result.Scan(&trait.Id, &trait.ParentId, &trait.Name, &trait.Description)
		helper.PanicIfError(err)

		traits = append(traits, trait)
	}
	return traits
}

// SelectById implements TraitRepository.
func (r *TraitRepositoryImpl) SelectById(ctx context.Context, traitId int) (model.Trait, error) {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from trait where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, traitId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	trait := model.Trait{}

	if result.Next() {
		err := result.Scan(&trait.Id, &trait.ParentId, &trait.Name, &trait.Description)
		helper.PanicIfError(err)
		return trait, nil
	} else {
		return trait, errors.New("trait with that id not found")
	}
}

// Update implements TraitRepository.
func (r *TraitRepositoryImpl) Update(ctx context.Context, trait model.Trait) model.Trait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update trait set parentid=$1,name=$2,description=$3 where id=$4"
	_, errExec := tx.ExecContext(ctx, SQL, trait.ParentId, trait.Name, trait.Description, trait.Id)
	helper.PanicIfError(errExec)
	return trait
}

// SelectByParentId implements TraitRepository.
func (r *TraitRepositoryImpl) SelectByParentId(ctx context.Context, parentId int) []model.Trait {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from trait where parentId=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, parentId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var traits []model.Trait

	for result.Next() {
		trait := model.Trait{}
		err := result.Scan(&trait.Id, &trait.ParentId, &trait.Name, &trait.Description)
		helper.PanicIfError(err)

		traits = append(traits, trait)
	}
	return traits
}

func NewTraitRepository(Db *sql.DB) TraitRepository {
	return &TraitRepositoryImpl{Db: Db}
}
