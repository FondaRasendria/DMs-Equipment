package request

type SubraceUpdateRequest struct {
	Id          int    "validate: 'required' json: 'id'"
	Description string "validate: 'required' json: 'description'"
	ParentId    int    "validate: 'required' json: 'parentid'"
	Source      string "validate: 'required' json: 'source'"
	Ability     string "validate: 'required' json: 'ability'"
	Age         string "validate: 'required' json: 'age'"
	Alignment   string "validate: 'required' json: 'alignment'"
	Size        string "validate: 'required' json: 'size'"
	Speed       string "validate: 'required' json: 'speed'"
	Language    string "validate: 'required' json: 'language'"
}
