package request

type MonsterConditionImmunUpdateRequest struct {
	MonsterId   int `json:"monster_id"`
	ConditionId int `json:"condition_id"`
}
