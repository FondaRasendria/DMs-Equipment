package model

type Monster struct {
	Id                int    `json:"id" gorm:"primaryKey"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Alignment         string `json:"alignment"`
	AC                int    `json:"ac"`
	FixedHP           int    `json:"fixed_hp"`
	DiceHP            string `json:"dice_hp"`
	Speed             int    `json:"speed"`
	STR               int    `json:"str"`
	DEX               int    `json:"dex"`
	CON               int    `json:"con"`
	INT               int    `json:"int"`
	WIS               int    `json:"wis"`
	CHA               int    `json:"cha"`
	Senses            string `json:"senses"`
	PassivePerception int    `json:"passive_perception"`
	CR                string `json:"cr"`
	XPReward          int    `json:"xp_reward"`
	Description       string `json:"description"`
	Source            string `json:"source"`
}
