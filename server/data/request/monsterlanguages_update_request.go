package request

type MonsterLanguagesUpdateRequest struct {
	MonsterId  int `json:"monster_id"`
	LanguageId int `json:"language_id"`
}
