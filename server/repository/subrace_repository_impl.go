package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud/helper"
	"golang-crud/model"
)

type SubraceRepositoryImpl struct {
	Db *sql.DB
}

// Create implements SubraceRepository.
func (r *SubraceRepositoryImpl) Create(ctx context.Context, subrace model.Subrace) model.Subrace {
	tx1, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx1)

	tx2, err2 := r.Db.Begin()
	helper.PanicIfError(err2)
	defer helper.CommitOrRollback(tx2)

	SQLid := "select max(id) from subrace"
	result, errQuery := tx1.QueryContext(ctx, SQLid)
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

	SQL := "insert into subrace(id, parentid, source, ability, age, alignment, size, speed, language, description) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	_, errExec := tx2.ExecContext(ctx, SQL, maxId, subrace.ParentId, subrace.Source, subrace.Ability, subrace.Age, subrace.Alignment, subrace.Size, subrace.Speed, subrace.Language, subrace.Description)
	helper.PanicIfError(errExec)
	return model.Subrace{Id: maxId, ParentId: subrace.ParentId, Source: subrace.Source, Ability: subrace.Ability, Age: subrace.Age, Alignment: subrace.Alignment, Size: subrace.Size, Speed: subrace.Speed, Language: subrace.Language, Description: subrace.Description}
}

// Delete implements SubraceRepository.
func (r *SubraceRepositoryImpl) Delete(ctx context.Context, subraceId int) model.Subrace {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	subrace, errSelect := r.SelectById(ctx, subraceId)
	helper.PanicIfError(errSelect)

	SQL := "delete from subrace where id=$1"
	_, errExec := tx.ExecContext(ctx, SQL, subraceId)
	helper.PanicIfError(errExec)
	return subrace
}

// SelectAll implements SubraceRepository.
func (r *SubraceRepositoryImpl) SelectAll(ctx context.Context) []model.Subrace {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from subrace"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var subraces []model.Subrace

	for result.Next() {
		subrace := model.Subrace{}
		err := result.Scan(&subrace.Id, &subrace.ParentId, &subrace.Source, &subrace.Ability, &subrace.Age, &subrace.Alignment, &subrace.Size, &subrace.Speed, &subrace.Language, &subrace.Description)
		helper.PanicIfError(err)

		subraces = append(subraces, subrace)
	}
	return subraces
}

// SelectById implements SubraceRepository.
func (r *SubraceRepositoryImpl) SelectById(ctx context.Context, subraceId int) (model.Subrace, error) {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from subrace where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, subraceId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	subrace := model.Subrace{}

	if result.Next() {
		err := result.Scan(&subrace.Id, &subrace.ParentId, &subrace.Source, &subrace.Ability, &subrace.Age, &subrace.Alignment, &subrace.Size, &subrace.Speed, &subrace.Language, &subrace.Description)
		helper.PanicIfError(err)
		return subrace, nil
	} else {
		return subrace, errors.New("subrace with that id not found")
	}
}

// Update implements SubraceRepository.
func (r *SubraceRepositoryImpl) Update(ctx context.Context, subrace model.Subrace) model.Subrace {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update subrace set parentid=$1,source=$2,ability=$3,age=$4,alignment=$5,size=$6,speed=$7,language=$8,description=$9 where id=$10"
	_, errExec := tx.ExecContext(ctx, SQL, subrace.ParentId, subrace.Description, subrace.Source, subrace.Ability, subrace.Age, subrace.Alignment, subrace.Size, subrace.Speed, subrace.Language, subrace.Description, subrace.Id)
	helper.PanicIfError(errExec)
	return subrace
}

// SelectByParentId implements SubraceRepository.
func (r *SubraceRepositoryImpl) SelectByParentId(ctx context.Context, parentId int) []model.Subrace {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from subrace where parentid=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, parentId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var subraces []model.Subrace

	for result.Next() {
		subrace := model.Subrace{}
		err := result.Scan(&subrace.Id, &subrace.ParentId, &subrace.Source, &subrace.Ability, &subrace.Age, &subrace.Alignment, &subrace.Size, &subrace.Speed, &subrace.Language, &subrace.Description)
		helper.PanicIfError(err)

		subraces = append(subraces, subrace)
	}
	return subraces
}

func NewSubraceRepository(Db *sql.DB) SubraceRepository {
	return &SubraceRepositoryImpl{Db: Db}
}
