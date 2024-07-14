package request

type SubraceCreateRequest struct {
	ParentId    int    `validate:"required" json:"parentid"`
	Description string `validate:"required" json:"description"`
	Source      string `validate:"required" json:"source"`
	Ability     string `json:"ability"`
	Age         string `json:"age"`
	Alignment   string `json:"alignment"`
	Size        string `json:"size"`
	Speed       string `json:"speed"`
	Language    string `json:"language"`
}
