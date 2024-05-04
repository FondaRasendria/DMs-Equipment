package request

type RaceCreateRequest struct {
	Name string "validate: 'required' json: 'name'"
}
