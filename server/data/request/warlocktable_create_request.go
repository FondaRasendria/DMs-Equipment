package request

type WarlockTableCreateRequest struct {
	ProfBonus       int    `json:"prof_bonus" validate:"required"`
	CantripKnown    int    `json:"cantrip_known" validate:"required"`
	SpellKnown      int    `json:"spell_known" validate:"required"`
	SpellSlot       int    `json:"spell_slot" validate:"required"`
	SlotLevel       string `json:"slot_level" validate:"required"`
	InvocationKnown int    `json:"invocation_known" validate:"required"`
}
