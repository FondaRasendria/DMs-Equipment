package response

type SubclassResponse struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parent_id"`
	Name        string `json:"name"`
	Source      string `json:"source"`
	Description string `json:"description"`
}
