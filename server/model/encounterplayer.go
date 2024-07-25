package model

type EncounterPlayer struct {
	Id          int `json:"id" gorm:"primarykey"`
	EncounterId int `json:"encounter_id"`
	PlayerId    int `json:"player_id"`
	Initiative  int `json:"initiative"`
	CurrentHP   int `json:"current_hp"`
}
