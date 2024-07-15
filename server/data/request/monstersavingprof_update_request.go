package request

type MonsterSavingProfUpdateRequest struct {
	MonsterId int `json:"monster_id"`
	SavingId  int `json:"saving_id"`
}
