package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud/helper"
	"golang-crud/model"
)

type MonsterRepositoryImpl struct {
	Db *sql.DB
}

// Create implements MonsterRepository.
func (r *MonsterRepositoryImpl) Create(ctx context.Context, monster model.Monster) model.Monster {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tx2, err2 := r.Db.Begin()
	helper.PanicIfError(err2)
	defer helper.CommitOrRollback(tx2)

	SQLid := "select max(id) from monster"
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

	SQL := "insert into monster(id, name, type, alignment, ac, fixedhp, dicehp, speed, str, dex, con, int, wis, cha, savingthrows, skills, damageimmunities, damagevulnerabilities, damageresistences, conditionimmunities, senses, language, cr, description, source) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)"
	_, errExec := tx2.ExecContext(ctx, SQL, maxId, monster.Name, monster.Type, monster.Alignment, monster.AC, monster.FixedHP, monster.DiceHP, monster.Speed, monster.STR, monster.DEX, monster.CON, monster.INT, monster.WIS, monster.CHA, monster.SavingThrows, monster.Skills, monster.DamageImmunities, monster.DamageVulnerabilities, monster.DamageResistences, monster.ConditionImmunities, monster.Senses, monster.Language, monster.CR, monster.Description, monster.Source)
	helper.PanicIfError(errExec)
	return model.Monster{Id: maxId, Name: monster.Name, Type: monster.Type, Alignment: monster.Alignment, AC: monster.AC, FixedHP: monster.FixedHP, DiceHP: monster.DiceHP, Speed: monster.Speed, STR: monster.STR, DEX: monster.DEX, CON: monster.CON, INT: monster.INT, WIS: monster.WIS, CHA: monster.CHA, SavingThrows: monster.SavingThrows, Skills: monster.Skills, DamageImmunities: monster.DamageImmunities, DamageVulnerabilities: monster.DamageVulnerabilities, DamageResistences: monster.DamageResistences, ConditionImmunities: monster.ConditionImmunities, Senses: monster.Senses, Language: monster.Language, CR: monster.CR, Description: monster.Description, Source: monster.Source}
}

// Delete implements MonsterRepository.
func (r *MonsterRepositoryImpl) Delete(ctx context.Context, monsterId int) model.Monster {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	monster, errSelect := r.SelectById(ctx, monsterId)
	helper.PanicIfError(errSelect)

	SQL := "delete from monster where id=$1"
	_, errExec := tx.ExecContext(ctx, SQL, monsterId)
	helper.PanicIfError(errExec)
	return monster
}

// SelectAll implements MonsterRepository.
func (r *MonsterRepositoryImpl) SelectAll(ctx context.Context) []model.Monster {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from monster"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var monsters []model.Monster

	for result.Next() {
		monster := model.Monster{}
		err := result.Scan(&monster.Id, &monster.Name, &monster.Type, &monster.Alignment, &monster.AC, &monster.FixedHP, &monster.DiceHP, &monster.Speed, &monster.STR, &monster.DEX, &monster.CON, &monster.INT, &monster.WIS, &monster.CHA, &monster.SavingThrows, &monster.Skills, &monster.DamageImmunities, &monster.DamageVulnerabilities, &monster.DamageResistences, &monster.ConditionImmunities, &monster.Senses, &monster.Language, &monster.CR, &monster.Description, &monster.Source)
		helper.PanicIfError(err)

		monsters = append(monsters, monster)
	}
	return monsters
}

// SelectById implements MonsterRepository.
func (r *MonsterRepositoryImpl) SelectById(ctx context.Context, monsterId int) (model.Monster, error) {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from monster where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, monsterId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	monster := model.Monster{}

	if result.Next() {
		err := result.Scan(&monster.Id, &monster.Name, &monster.Type, &monster.Alignment, &monster.AC, &monster.FixedHP, &monster.DiceHP, &monster.Speed, &monster.STR, &monster.DEX, &monster.CON, &monster.INT, &monster.WIS, &monster.CHA, &monster.SavingThrows, &monster.Skills, &monster.DamageImmunities, &monster.DamageVulnerabilities, &monster.DamageResistences, &monster.ConditionImmunities, &monster.Senses, &monster.Language, &monster.CR, &monster.Description, &monster.Source)
		helper.PanicIfError(err)
		return monster, nil
	} else {
		return monster, errors.New("race with that id not found")
	}
}

// SelectByName implements MonsterRepository.
func (r *MonsterRepositoryImpl) SelectByName(ctx context.Context, name string) (model.Monster, error) {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from monster where name=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(errQuery)
	defer result.Close()

	monster := model.Monster{}

	if result.Next() {
		err := result.Scan(&monster.Id, &monster.Name, &monster.Type, &monster.Alignment, &monster.AC, &monster.FixedHP, &monster.DiceHP, &monster.Speed, &monster.STR, &monster.DEX, &monster.CON, &monster.INT, &monster.WIS, &monster.CHA, &monster.SavingThrows, &monster.Skills, &monster.DamageImmunities, &monster.DamageVulnerabilities, &monster.DamageResistences, &monster.ConditionImmunities, &monster.Senses, &monster.Language, &monster.CR, &monster.Description, &monster.Source)
		helper.PanicIfError(err)
		return monster, nil
	} else {
		return monster, errors.New("race with that name not found")
	}
}

// SelectBySource implements MonsterRepository.
func (r *MonsterRepositoryImpl) SelectBySource(ctx context.Context, source string) []model.Monster {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select * from monster where source=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, source)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var monsters []model.Monster

	for result.Next() {
		monster := model.Monster{}
		err := result.Scan(&monster.Id, &monster.Name, &monster.Type, &monster.Alignment, &monster.AC, &monster.FixedHP, &monster.DiceHP, &monster.Speed, &monster.STR, &monster.DEX, &monster.CON, &monster.INT, &monster.WIS, &monster.CHA, &monster.SavingThrows, &monster.Skills, &monster.DamageImmunities, &monster.DamageVulnerabilities, &monster.DamageResistences, &monster.ConditionImmunities, &monster.Senses, &monster.Language, &monster.CR, &monster.Description, &monster.Source)
		helper.PanicIfError(err)

		monsters = append(monsters, monster)
	}
	return monsters
}

// Update implements MonsterRepository.
func (r *MonsterRepositoryImpl) Update(ctx context.Context, monster model.Monster) model.Monster {
	tx, err := r.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update monster set name=$1,type=$2,alignment=$3,ac=$4,fixedhp=$5,dicehp=$6,speed=$7,str=$8,dex=$9,con=$10,int=$11,wis=$12,cha=$13,savingthrows=$14,skills=$15,damageimmunities=$16,damagevulnerabilities=$17,damageresistences=$18,conditionimmunities=$19,senses=$20,language=$21,cr=$22,description=$23,source=$24 where id=$25"
	_, errExec := tx.ExecContext(ctx, SQL, monster.Name, monster.Type, monster.Alignment, monster.AC, monster.FixedHP, monster.DiceHP, monster.Speed, monster.STR, monster.DEX, monster.CON, monster.INT, monster.WIS, monster.CHA, monster.SavingThrows, monster.Skills, monster.DamageImmunities, monster.DamageVulnerabilities, monster.DamageResistences, monster.ConditionImmunities, monster.Senses, monster.Language, monster.CR, monster.Description, monster.Source, monster.Id)
	helper.PanicIfError(errExec)
	return monster
}

func NewMonsterRepository(Db *sql.DB) MonsterRepository {
	return &MonsterRepositoryImpl{Db: Db}
}
