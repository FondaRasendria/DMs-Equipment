package request

type BardTableCreateRequest struct {
	ProfBonus    int `json:"prof_bonus" validate:"required"`
	CantripKnown int `json:"cantrip_known" validate:"required"`
	SpellKnown   int `json:"spell_known" validate:"required"`
	SpellSlot1   int `json:"spell_slot1" validate:"required"`
	SpellSlot2   int `json:"spell_slot2" validate:"required"`
	SpellSlot3   int `json:"spell_slot3" validate:"required"`
	SpellSlot4   int `json:"spell_slot4" validate:"required"`
	SpellSlot5   int `json:"spell_slot5" validate:"required"`
	SpellSlot6   int `json:"spell_slot6" validate:"required"`
	SpellSlot7   int `json:"spell_slot7" validate:"required"`
	SpellSlot8   int `json:"spell_slot8" validate:"required"`
	SpellSlot9   int `json:"spell_slot9" validate:"required"`
}
