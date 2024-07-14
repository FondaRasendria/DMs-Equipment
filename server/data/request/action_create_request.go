package request

type ActionCreateRequest struct {
	ParentId    int    `json:"parent_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
