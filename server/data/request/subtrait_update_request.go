package request

type SubtraitUpdateRequest struct {
	Id          int    "validate: 'required' json: 'id'"
	ParentId    int    "validate: 'required' json: 'parentid'"
	Level       int    "validate: 'required' json: 'level'"
	Description string "validate: 'required' json: 'description'"
}
