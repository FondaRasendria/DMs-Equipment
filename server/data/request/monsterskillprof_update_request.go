package request

type MonsterSkillProfUpdateRequest struct {
	MonsterId int `json:"monster_id"`
	SkillId   int `json:"skill_id"`
}
