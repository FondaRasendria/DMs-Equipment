package request

type MonsterSkillProfCreateRequest struct {
	MonsterId int `validate:"required" json:"monster_id"`
	SkillId   int `validate:"required" json:"skill_id"`
}
