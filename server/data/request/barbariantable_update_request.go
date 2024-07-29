package request

type BarbarianTableUpdateRequest struct {
	ProfBonus  int `json:"prof_bonus"`
	Rage       int `json:"rage"`
	RageDamage int `json:"rage_damage"`
}
