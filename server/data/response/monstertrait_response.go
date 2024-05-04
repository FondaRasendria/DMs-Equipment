package response

type MonsterTraitResponse struct {
	Id          int    "json: 'id'"
	ParentId    int    "json: 'parentid'"
	Name        string "json: 'name'"
	Description string "json: 'description'"
}
