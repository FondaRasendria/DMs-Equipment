package model

type WarlockTable struct {
	Id              int    `json:"id" gorm:"primarykey"`
	ProfBonus       int    `json:"prof_bonus"`
	CantripKnown    int    `json:"cantrip_known"`
	SpellKnown      int    `json:"spell_known"`
	SpellSlot       int    `json:"spell_slot"`
	SlotLevel       string `json:"slot_level"`
	InvocationKnown int    `json:"invocation_known"`
}
