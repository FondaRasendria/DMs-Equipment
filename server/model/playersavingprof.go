package model

type PlayerSavingProf struct {
	Id       int `json:"id" gorm:"primarykey"`
	PlayerId int `json:"player_id"`
	SavingId int `json:"saving_id"`
}
