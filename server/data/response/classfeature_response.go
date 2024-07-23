package response

type ClassFeatureResponse struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parent_id"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}
