package model

type SavingThrow struct {
	Id   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}
