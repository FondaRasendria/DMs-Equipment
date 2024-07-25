package request

type PlayerSkillUpdateRequest struct {
	PlayerId int `json:"player_id"`
	SkillId  int `json:"skill_id"`
}
