package request

type TraitCreateRequest struct {
	ParentId    int    "validate: 'required' json: 'parentid'"
	Name        string "validate: 'required' json: 'name'"
	Description string "validate: 'required' json: 'description'"
}
