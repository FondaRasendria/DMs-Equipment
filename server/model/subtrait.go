package model

type Subtrait struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	ParentId    int    `json:"parent_id"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}
