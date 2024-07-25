package model

type PlayerConditionImmunities struct {
	Id          int `json:"id" gorm:"primarykey"`
	PlayerId    int `json:"player_id"`
	ConditionId int `json:"condition_id"`
}
