package response

type EncounterPlayerResponse struct {
	Id          int `json:"id"`
	EncounterId int `json:"encounter_id"`
	PlayerId    int `json:"player_id"`
	Initiative  int `json:"initiative"`
	CurrentHP   int `json:"current_hp"`
}
