package response

type BackgroundResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Source      string `json:"source"`
	SkillProf   string `json:"skillprof"`
	ToolProf    string `json:"toolprof"`
	Language    string `json:"language"`
	Equipment   string `json:"equipment"`
}
