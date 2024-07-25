package response

type PlayerLanguageResponse struct {
	Id         int `json:"id"`
	PlayerId   int `json:"player_id"`
	LanguageId int `json:"language_id"`
}
