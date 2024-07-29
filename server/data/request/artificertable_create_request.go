package request

type ArtificerTableCreateRequest struct {
	ProfBonus     int `json:"prof_bonus" validate:"required"`
	InfusionKnown int `json:"infusion_known" validate:"required"`
	InfusedItem   int `json:"infused_item" validate:"required"`
	CantripKnown  int `json:"cantrip_known" validate:"required"`
	SpellSlot1    int `json:"spell_slot1" validate:"required"`
	SpellSlot2    int `json:"spell_slot2" validate:"required"`
	SpellSlot3    int `json:"spell_slot3" validate:"required"`
	SpellSlot4    int `json:"spell_slot4" validate:"required"`
	SpellSlot5    int `json:"spell_slot5" validate:"required"`
}
