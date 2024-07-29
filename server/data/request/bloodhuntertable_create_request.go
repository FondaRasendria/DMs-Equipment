package request

type BloodHunterTableCreateRequest struct {
	ProfBonus       int    `json:"prof_bonus" validate:"required"`
	HemocraftDie    string `json:"hemocraft_die" validate:"required"`
	BloodCurseKnown int    `json:"blood_curse_known" validate:"required"`
}
