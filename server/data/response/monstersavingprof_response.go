package response

type MonsterSavingProfResponse struct {
	Id        int `json:"id"`
	MonsterId int `json:"monster_id"`
	SavingId  int `json:"saving_id"`
}
