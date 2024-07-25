package request

type PlayerSavingUpdateRequest struct {
	PlayerId int `json:"player_id"`
	SavingId int `json:"saving_id"`
}
