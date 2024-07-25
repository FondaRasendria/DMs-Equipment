package response

type PlayerSkillResponse struct {
	Id       int `json:"id"`
	PlayerId int `json:"player_id"`
	SkillId  int `json:"skill_id"`
}
