package request

type SubclassCreateRequest struct {
	ParentId    int    `json:"parent_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Source      string `json:"source"`
	Description string `json:"description" validate:"required"`
}
