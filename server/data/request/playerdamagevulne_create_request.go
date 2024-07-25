package request

type PlayerDamageVulneCreateRequest struct {
	PlayerId int `json:"player_id" validate:"required"`
	DamageId int `json:"damage_id" validate:"required"`
}
