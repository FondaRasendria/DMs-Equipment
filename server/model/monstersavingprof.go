package model

type MonsterSavingProf struct {
	Id        int `json:"id" gorm:"primarykey"`
	MonsterId int `json:"monster_id"`
	SavingId  int `json:"saving_id"`
}
