package request

type EncounterUpdateRequest struct {
	CampaignId int    `json:"campaign_id"`
	Name       string `json:"name"`
	Turn       int    `json:"turn"`
	Round      int    `json:"round"`
}
