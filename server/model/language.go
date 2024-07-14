package model

type Language struct {
	Id          int    `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Alphabet    string `json:"alphabet"`
	MainSpeaker string `json:"main_speaker"`
}
