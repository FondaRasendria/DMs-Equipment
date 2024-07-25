package request

type PlayerDamageImmunCreateRequest struct {
	PlayerId int `json:"player_id" validate:"required"`
	DamageId int `json:"damage_id" validate:"required"`
}
