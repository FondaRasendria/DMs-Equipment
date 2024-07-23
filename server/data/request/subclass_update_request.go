package request

type SubclassUpdateRequest struct {
	ParentId    int    `json:"parent_id"`
	Name        string `json:"name"`
	Source      string `json:"source"`
	Description string `json:"description"`
}
