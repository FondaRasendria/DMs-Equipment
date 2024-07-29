package request

type BloodHunterTableUpdateRequest struct {
	ProfBonus       int    `json:"prof_bonus"`
	HemocraftDie    string `json:"hemocraft_die"`
	BloodCurseKnown int    `json:"blood_curse_known"`
}
