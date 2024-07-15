package model

type MonsterLanguages struct {
	Id         int `json:"id" gorm:"primarykey"`
	MonsterId  int `json:"monster_id"`
	LanguageId int `json:"language_id"`
}
