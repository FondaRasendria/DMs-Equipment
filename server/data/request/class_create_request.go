package request

type ClassCreateRequest struct {
	Name           string `json:"name" validation:"required"`
	Description    string `json:"description" validation:"required"`
	Multiclass     string `json:"multiclass"`
	TableId        int    `json:"table_id" validation:"required"`
	HitDice        string `json:"hit_dice" validation:"required"`
	HitPointFirst  string `json:"hit_point_first" validation:"required"`
	HitPointHigher string `json:"hit_point_higher" validation:"required"`
	ArmorProf      string `json:"armor_prof"`
	WeaponProf     string `json:"weapon_prof"`
	ToolProf       string `json:"tool_prof"`
	SavingProf     string `json:"saving_prof"`
	SkillProf      string `json:"skill_prof"`
	Equipment      string `json:"equipment"`
}
