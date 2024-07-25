package model

type PlayerDamageImmunities struct {
	Id       int `json:"id" gorm:"primarykey"`
	PlayerId int `json:"player_id"`
	DamageId int `json:"damage_id"`
}
