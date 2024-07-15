package request

type MonsterDmgVulneCreateRequest struct {
	MonsterId int `validate:"required" json:"monster_id"`
	DamageId  int `validate:"required" json:"damage_id"`
}
