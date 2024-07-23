package response

type ClassResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Multiclass     string `json:"multiclass"`
	TableId        int    `json:"table_id"`
	HitDice        string `json:"hit_dice"`
	HitPointFirst  string `json:"hit_point_first"`
	HitPointHigher string `json:"hit_point_higher"`
	ArmorProf      string `json:"armor_prof"`
	WeaponProf     string `json:"weapon_prof"`
	ToolProf       string `json:"tool_prof"`
	SavingProf     string `json:"saving_prof"`
	SkillProf      string `json:"skill_prof"`
	Equipment      string `json:"equipment"`
}
