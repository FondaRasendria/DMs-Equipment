package model

type MonsterTrait struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	ParentId    int    `json:"parent_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
