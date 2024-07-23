package request

type ClassFeatureCreateRequest struct {
	ParentId    int    `json:"parent_id" validate:"required"`
	Level       int    `json:"level" validate:"required"`
	Description string `json:"description" validate:"required"`
}
