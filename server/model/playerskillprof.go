package model

type PlayerSkillProf struct {
	Id       int `json:"id" gorm:"primarykey"`
	PlayerId int `json:"player_id"`
	SkillId  int `json:"skill_id"`
}
