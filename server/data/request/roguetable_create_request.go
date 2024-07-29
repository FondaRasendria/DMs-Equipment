package request

type RogueTableCreateRequest struct {
	ProfBonus   int    `json:"prof_bonus" validate:"required"`
	SneakAttack string `json:"sneak_attack" validate:"required"`
}
