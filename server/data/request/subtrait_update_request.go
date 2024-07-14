package request

type SubtraitUpdateRequest struct {
	ParentId    int    `json:"parentid"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}
