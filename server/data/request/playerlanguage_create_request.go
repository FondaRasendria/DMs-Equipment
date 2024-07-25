package request

type PlayerLanguageCreateRequest struct {
	PlayerId   int `json:"player_id" validate:"required"`
	LanguageId int `json:"language_id" validate:"required"`
}
