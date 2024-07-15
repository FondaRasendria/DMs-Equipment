package request

type MonsterDmgResisUpdateRequest struct {
	MonsterId int `json:"monster_id"`
	DamageId  int `json:"damage_id"`
}
