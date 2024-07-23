package model

type Subclass struct {
	Id          int    `json:"id" gorm:"primarykey"`
	ParentId    int    `json:"parent_id"`
	Name        string `json:"name"`
	Source      string `json:"source"`
	Description string `json:"description"`
}
