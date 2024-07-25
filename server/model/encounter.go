package model

type Encounter struct {
	Id         int    `json:"id" gorm:"primarykey"`
	CampaignId int    `json:"campaign_id"`
	Name       string `json:"name"`
	Turn       int    `json:"turn"`
	Round      int    `json:"round"`
}
