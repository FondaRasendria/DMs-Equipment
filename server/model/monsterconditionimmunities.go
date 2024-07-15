package model

type MonsterConditionImmunities struct {
	Id          int `json:"id" gorm:"primarykey"`
	MonsterId   int `json:"monster_id"`
	ConditionId int `json:"condition_id"`
}
