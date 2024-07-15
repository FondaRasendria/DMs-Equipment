package response

type MonsterSkillProfResponse struct {
	Id        int `json:"id"`
	MonsterId int `json:"monster_id"`
	SkillId   int `json:"skill_id"`
}
