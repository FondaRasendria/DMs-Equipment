package response

type PlayerDamageVulneResponse struct {
	Id       int `json:"id"`
	PlayerId int `json:"player_id"`
	DamageId int `json:"damage_id"`
}
