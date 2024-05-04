package request

type MonsterTraitUpdateRequest struct {
	Id          int    "validate: 'required' json: 'id'"
	ParentId    int    "validate: 'required' json: 'parentid'"
	Name        string "validate: 'required' json: 'name'"
	Description string "validate: 'required' json: 'description'"
}
