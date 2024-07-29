package request

type BarbarianTableCreateRequest struct {
	ProfBonus  int `json:"prof_bonus" validate:"required"`
	Rage       int `json:"rage" validate:"required"`
	RageDamage int `json:"rage_damage" validate:"required"`
}
