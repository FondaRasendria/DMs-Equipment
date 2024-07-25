package request

type EncounterMonsterCreateRequest struct {
	EncounterId int `json:"encounter_id"`
	MonsterId   int `json:"monster_id"`
	Initiative  int `json:"initiative"`
	CurrentHP   int `json:"current_hp"`
}
