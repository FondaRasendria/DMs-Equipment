package request

type PlayerConditionImmunCreateRequest struct {
	PlayerId    int `json:"player_id" validate:"required"`
	ConditionId int `json:"condition_id" validate:"required"`
}
