package request

type SubraceUpdateRequest struct {
	Description string `json:"description"`
	ParentId    int    `json:"parentid"`
	Source      string `json:"source"`
	Ability     string `json:"ability"`
	Age         string `json:"age"`
	Alignment   string `json:"alignment"`
	Size        string `json:"size"`
	Speed       string `json:"speed"`
	Language    string `json:"language"`
}
