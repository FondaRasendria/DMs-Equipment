package response

type PlayerConditionImmunResponse struct {
	Id          int `json:"id"`
	PlayerId    int `json:"player_id"`
	ConditionId int `json:"condition_id"`
}
