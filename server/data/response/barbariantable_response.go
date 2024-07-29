package response

type BarbarianTableResponse struct {
	Id         int `json:"id"`
	ProfBonus  int `json:"prof_bonus"`
	Rage       int `json:"rage"`
	RageDamage int `json:"rage_damage"`
}
