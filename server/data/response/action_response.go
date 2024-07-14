package response

type ActionResponse struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parent_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
