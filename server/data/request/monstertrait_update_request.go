package request

type MonsterTraitUpdateRequest struct {
	ParentId    int    `json:"parentid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
