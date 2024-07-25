package request

type PlayerCreateRequest struct {
	PlayerName        string `json:"player_name"`
	CharacterName     string `json:"character_name" validate:"required"`
	RaceId            int    `json:"race_id" validate:"required"`
	ClassId           int    `json:"class_id" validate:"required"`
	SubclassId        int    `json:"subclass_id" validate:"required"`
	BackgroundId      int    `json:"background_id" validate:"required"`
	XP                int    `json:"xp" validate:"required"`
	Level             int    `json:"level" validate:"required"`
	AC                int    `json:"ac" validate:"required"`
	MaxHP             int    `json:"max_hp" validate:"required"`
	HitDice           string `json:"hit_dice"`
	SpellSaveDC       int    `json:"spell_save_dc"`
	SpellBonus        int    `json:"spell_bonus"`
	Speed             int    `json:"speed" validate:"required"`
	STR               int    `json:"str" validate:"required"`
	DEX               int    `json:"dex" validate:"required"`
	COS               int    `json:"cos" validate:"required"`
	INT               int    `json:"int" validate:"required"`
	WIS               int    `json:"wis" validate:"required"`
	CHA               int    `json:"cha" validate:"required"`
	PassivePerception int    `json:"passive_perception"`
	PassiveInsight    int    `json:"passive_insight"`
}
