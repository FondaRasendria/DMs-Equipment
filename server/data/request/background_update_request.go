package request

type BackgroundUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Source      string `json:"source"`
	SkillProf   string `json:"skillprof"`
	ToolProf    string `json:"toolprof"`
	Language    string `json:"language"`
	Equipment   string `json:"equipment"`
}
