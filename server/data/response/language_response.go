package response

type LanguageResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Alphabet    string `json:"alphabet"`
	MainSpeaker string `json:"main_speaker"`
}
