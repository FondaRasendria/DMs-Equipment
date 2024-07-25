package request

type PlayerLanguageUpdateRequest struct {
	PlayerId   int `json:"player_id"`
	LanguageId int `json:"language_id"`
}
