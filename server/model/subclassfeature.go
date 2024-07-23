package model

type SubclassFeature struct {
	Id          int    `json:"id" gorm:"primarykey"`
	ParentId    int    `json:"parent_id"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}
