package request

type MonsterSavingProfCreateRequest struct {
	MonsterId int `validate:"required" json:"monster_id"`
	SavingId  int `validate:"required" json:"saving_id"`
}
