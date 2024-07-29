package request

type FighterTableCreateRequest struct {
	ProfBonus int `json:"prof_bonus" validate:"required"`
}
