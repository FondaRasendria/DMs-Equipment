package request

type SubtraitCreateRequest struct {
	ParentId    int    `validate:"required" json:"parentid"`
	Level       int    `validate:"required" json:"level"`
	Description string `validate:"required" json:"description"`
}
