package model

type Race struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
