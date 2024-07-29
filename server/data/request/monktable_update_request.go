package request

type MonkTableUpdateRequest struct {
	ProfBonus       int    `json:"prof_bonus"`
	MartialArt      string `json:"martial_art"`
	KiPoint         int    `json:"ki_point"`
	UnarmedMovement string `json:"unarmed_movement"`
}
