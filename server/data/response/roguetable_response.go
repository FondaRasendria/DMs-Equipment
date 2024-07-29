package response

type RogueTableResponse struct {
	Id          int    `json:"id"`
	ProfBonus   int    `json:"prof_bonus"`
	SneakAttack string `json:"sneak_attack"`
}
