package request

type PlayerConditionImmunUpdateRequest struct {
	PlayerId    int `json:"player_id"`
	ConditionId int `json:"condition_id"`
}
