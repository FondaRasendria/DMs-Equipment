package request

type MonsterConditionImmunCreateRequest struct {
	MonsterId   int `validate:"required" json:"monster_id"`
	ConditionId int `validate:"required" json:"condition_id"`
}
