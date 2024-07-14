package model

type DamageType struct {
	Id   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}
