package model

type EncounterMonster struct {
	Id          int `json:"id" gorm:"primarykey"`
	EncounterId int `json:"encounter_id"`
	MonsterId   int `json:"monster_id"`
	Initiative  int `json:"initiative"`
	CurrentHP   int `json:"current_hp"`
}
