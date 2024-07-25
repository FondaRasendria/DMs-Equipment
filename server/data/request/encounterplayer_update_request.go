package request

type EncounterPlayerUpdateRequest struct {
	EncounterId int `json:"encounter_id"`
	PlayerId    int `json:"player_id"`
	Initiative  int `json:"initiative"`
	CurrentHP   int `json:"current_hp"`
}
