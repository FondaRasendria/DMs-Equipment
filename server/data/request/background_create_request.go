package request

type BackgroundCreateRequest struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
	Source      string `json:"source"`
	SkillProf   string `json:"skillprof"`
	ToolProf    string `json:"toolprof"`
	Language    string `json:"language"`
	Equipment   string `json:"equipment"`
}
