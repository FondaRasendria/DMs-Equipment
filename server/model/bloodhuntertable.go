package model

type BloodHunterTable struct {
	Id              int    `json:"id" gorm:"primarykey"`
	ProfBonus       int    `json:"prof_bonus"`
	HemocraftDie    string `json:"hemocraft_die"`
	BloodCurseKnown int    `json:"blood_curse_known"`
}
