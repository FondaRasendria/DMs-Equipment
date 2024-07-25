package request

type PlayerDamageVulneUpdateRequest struct {
	PlayerId int `json:"player_id"`
	DamageId int `json:"damage_id"`
}
