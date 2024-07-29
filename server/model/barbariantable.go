package model

type BarbarianTable struct {
	Id         int `json:"id" gorm:"primarykey"`
	ProfBonus  int `json:"prof_bonus"`
	Rage       int `json:"rage"`
	RageDamage int `json:"rage_damage"`
}
