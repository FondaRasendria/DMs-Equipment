package response

type SubraceResponse struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parentid"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Ability     string `json:"ability"`
	Age         string `json:"age"`
	Alignment   string `json:"alignment"`
	Size        string `json:"size"`
	Speed       string `json:"speed"`
	Language    string `json:"language"`
}
