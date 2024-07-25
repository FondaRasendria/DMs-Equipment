package response

type PlayerSavingResponse struct {
	Id       int `json:"id"`
	PlayerId int `json:"player_id"`
	SavingId int `json:"saving_id"`
}
