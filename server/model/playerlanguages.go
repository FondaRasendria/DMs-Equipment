package model

type PlayerLanguages struct {
	Id         int `json:"id" gorm:"primarykey"`
	PlayerId   int `json:"player_id"`
	LanguageId int `json:"language_id"`
}
