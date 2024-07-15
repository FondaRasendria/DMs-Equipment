package response

type MonsterDmgImmunResponse struct {
	Id        int `json:"id"`
	MonsterId int `json:"monster_id"`
	DamageId  int `json:"damage_id"`
}
