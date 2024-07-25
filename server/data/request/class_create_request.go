package request

type ClassCreateRequest struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description" validate:"required"`
	Multiclass     string `json:"multiclass"`
	TableId        int    `json:"table_id" validate:"required"`
	HitDice        string `json:"hit_dice" validate:"required"`
	HitPointFirst  string `json:"hit_point_first" validate:"required"`
	HitPointHigher string `json:"hit_point_higher" validate:"required"`
	ArmorProf      string `json:"armor_prof"`
	WeaponProf     string `json:"weapon_prof"`
	ToolProf       string `json:"tool_prof"`
	SavingProf     string `json:"saving_prof"`
	SkillProf      string `json:"skill_prof"`
	Equipment      string `json:"equipment"`
}
