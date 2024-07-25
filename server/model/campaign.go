package model

type Campaign struct {
	Id    int    `json:"id" gorm:"primarykey"`
	Title string `json:"title"`
	Notes string `json:"notes"`
}
