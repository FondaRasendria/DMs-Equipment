package request

type SpellCreateRequest struct {
	Name        string `validate:"required" json:"name"`
	Source      string `json:"source"`
	School      string `json:"school"`
	Level       int    `validate:"required" json:"level"`
	CastingTime string `json:"casting_time"`
	Range       string `json:"range"`
	Components  string `json:"components"`
	Duration    string `json:"duration"`
	Description string `validate:"required" json:"description"`
	HigherLevel string `json:"higher_level"`
}
