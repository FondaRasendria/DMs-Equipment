package response

type PlayerResponse struct {
	Id                int    `json:"id"`
	PlayerName        string `json:"player_name"`
	CharacterName     string `json:"character_name"`
	RaceId            int    `json:"race_id"`
	ClassId           int    `json:"class_id"`
	SubclassId        int    `json:"subclass_id"`
	BackgroundId      int    `json:"background_id"`
	XP                int    `json:"xp"`
	Level             int    `json:"level"`
	AC                int    `json:"ac"`
	MaxHP             int    `json:"max_hp"`
	HitDice           string `json:"hit_dice"`
	SpellSaveDC       int    `json:"spell_save_dc"`
	SpellBonus        int    `json:"spell_bonus"`
	Speed             int    `json:"speed"`
	STR               int    `json:"str"`
	DEX               int    `json:"dex"`
	COS               int    `json:"cos"`
	INT               int    `json:"int"`
	WIS               int    `json:"wis"`
	CHA               int    `json:"cha"`
	PassivePerception int    `json:"passive_perception"`
	PassiveInsight    int    `json:"passive_insight"`
}
