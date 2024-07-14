package request

type LanguageCreateRequest struct {
	Name        string `validate:"required" json:"name"`
	Alphabet    string `validate:"required" json:"alphabet"`
	MainSpeaker string `validate:"required" json:"main_speaker"`
}
