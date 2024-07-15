package model

type MonsterDamageImmunities struct {
	Id        int `json:"id" gorm:"primarykey"`
	MonsterId int `json:"monster_id"`
	DamageId  int `json:"damage_id"`
}
