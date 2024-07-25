package request

type PlayerSavingCreateRequest struct {
	PlayerId int `json:"player_id" validate:"required"`
	SavingId int `json:"saving_id" validate:"required"`
}
