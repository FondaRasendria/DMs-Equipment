package model

type Spell struct {
	Id          int    `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Source      string `json:"source"`
	School      string `json:"school"`
	Level       int    `json:"level"`
	CastingTime string `json:"casting_time"`
	Range       string `json:"range"`
	Components  string `json:"components"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
	HigherLevel string `json:"higher_level"`
}
