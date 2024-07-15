package request

type MonsterLanguagesCreateRequest struct {
	MonsterId  int `validate:"required" json:"monster_id"`
	LanguageId int `validate:"required" json:"language_id"`
}
