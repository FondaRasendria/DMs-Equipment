package model

type SkillType struct {
	Id   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}
