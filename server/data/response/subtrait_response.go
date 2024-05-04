package response

type SubtraitResponse struct {
	Id          int    "json: 'id'"
	ParentId    int    "json: 'parentid'"
	Level       int    "json: 'level'"
	Description string "json: 'description'"
}
