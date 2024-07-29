package request

type PaladinTableCreateRequest struct {
	ProfBonus  int `json:"prof_bonus" validate:"required"`
	SpellSlot1 int `json:"spell_slot1" validate:"required"`
	SpellSlot2 int `json:"spell_slot2" validate:"required"`
	SpellSlot3 int `json:"spell_slot3" validate:"required"`
	SpellSlot4 int `json:"spell_slot4" validate:"required"`
	SpellSlot5 int `json:"spell_slot5" validate:"required"`
}
