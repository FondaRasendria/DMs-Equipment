package response

type MonsterConditionImmunResponse struct {
	Id          int `json:"id"`
	MonsterId   int `json:"monster_id"`
	ConditionId int `json:"condition_id"`
}
