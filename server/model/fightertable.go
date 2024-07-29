package model

type FighterTable struct {
	Id        int `json:"id" gorm:"primarykey"`
	ProfBonus int `json:"prof_bonus"`
}
