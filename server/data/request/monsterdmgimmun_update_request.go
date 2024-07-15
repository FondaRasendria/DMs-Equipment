package request

type MonsterDmgImmunUpdateRequest struct {
	MonsterId int `json:"monster_id"`
	DamageId  int `json:"damage_id"`
}
