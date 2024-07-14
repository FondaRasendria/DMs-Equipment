package model

type ConditionType struct {
	Id   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}
