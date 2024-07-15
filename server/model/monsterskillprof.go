package model

type MonsterSkillProf struct {
	Id        int `json:"id" gorm:"primarykey"`
	MonsterId int `json:"monster_id"`
	SkillId   int `json:"skill_id"`
}
