package request

type LanguageUpdateRequest struct {
	Name        string `json:"name"`
	Alphabet    string `json:"alphabet"`
	MainSpeaker string `json:"main_speaker"`
}
