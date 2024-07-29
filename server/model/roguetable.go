package model

type RogueTable struct {
	Id          int    `json:"id" gorm:"primarykey"`
	ProfBonus   int    `json:"prof_bonus"`
	SneakAttack string `json:"sneak_attack"`
}
