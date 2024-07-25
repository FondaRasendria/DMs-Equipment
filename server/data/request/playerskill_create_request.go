package request

type PlayerSkillCreateRequest struct {
	PlayerId int `json:"player_id" validate:"required"`
	SkillId  int `json:"skill_id" validate:"required"`
}
