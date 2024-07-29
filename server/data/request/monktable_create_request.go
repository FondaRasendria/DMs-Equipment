package request

type MonkTableCreateRequest struct {
	ProfBonus       int    `json:"prof_bonus" validate:"required"`
	MartialArt      string `json:"martial_art" validate:"required"`
	KiPoint         int    `json:"ki_point" validate:"required"`
	UnarmedMovement string `json:"unarmed_movement" validate:"required"`
}
