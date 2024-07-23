package request

type ClassFeatureUpdateRequest struct {
	ParentId    int    `json:"parent_id"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}
