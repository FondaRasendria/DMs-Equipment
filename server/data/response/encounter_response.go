package response

type EncounterResponse struct {
	Id         int    `json:"id"`
	CampaignId int    `json:"campaign_id"`
	Name       string `json:"name"`
	Turn       int    `json:"turn"`
	Round      int    `json:"round"`
}
