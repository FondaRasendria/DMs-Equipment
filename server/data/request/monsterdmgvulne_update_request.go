package request

type MonsterDmgVulneUpdateRequest struct {
	MonsterId int `json:"monster_id"`
	DamageId  int `json:"damage_id"`
}
