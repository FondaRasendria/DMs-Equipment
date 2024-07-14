package request

type SkillTypeCreateRequest struct {
	Name string `validate:"required" json:"name"`
}
