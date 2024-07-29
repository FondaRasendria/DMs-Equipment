package response

type BloodHunterTableResponse struct {
	Id              int    `json:"id"`
	ProfBonus       int    `json:"prof_bonus"`
	HemocraftDie    string `json:"hemocraft_die"`
	BloodCurseKnown int    `json:"blood_curse_known"`
}
