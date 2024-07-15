package response

type MonsterDmgResisResponse struct {
	Id        int `json:"id"`
	MonsterId int `json:"monster_id"`
	DamageId  int `json:"damage_id"`
}
