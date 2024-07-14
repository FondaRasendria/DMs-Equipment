package request

type ActionUpdateRequest struct {
	ParentId    int    `json:"parent_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
