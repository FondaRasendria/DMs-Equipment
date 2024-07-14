package response

type TraitResponse struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parentid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
