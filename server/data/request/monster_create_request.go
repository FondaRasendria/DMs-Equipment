package request

type MonsterCreateRequest struct {
	Name              string `validate:"required" json:"name"`
	Type              string `validate:"required" json:"type"`
	Alignment         string `validate:"required" json:"alignment"`
	AC                int    `validate:"required" json:"ac"`
	FixedHP           int    `validate:"required" json:"fixedhp"`
	DiceHP            string `validate:"required" json:"dicehp"`
	Speed             int    `validate:"required" json:"speed"`
	STR               int    `validate:"required" json:"str"`
	DEX               int    `validate:"required" json:"dex"`
	CON               int    `validate:"required" json:"con"`
	INT               int    `validate:"required" json:"int"`
	WIS               int    `validate:"required" json:"wis"`
	CHA               int    `validate:"required" json:"cha"`
	Senses            string `json:"senses"`
	PassivePerception int    `json:"passive_perception"`
	CR                string `validate:"required" json:"cr"`
	XPReward          int    `json:"xp_reward"`
	Description       string `validate:"required" json:"description"`
	Source            string `validate:"required" json:"source"`
}
