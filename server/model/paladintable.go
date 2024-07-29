package model

type PaladinTable struct {
	Id         int `json:"id" gorm:"primarykey"`
	ProfBonus  int `json:"prof_bonus"`
	SpellSlot1 int `json:"spell_slot1"`
	SpellSlot2 int `json:"spell_slot2"`
	SpellSlot3 int `json:"spell_slot3"`
	SpellSlot4 int `json:"spell_slot4"`
	SpellSlot5 int `json:"spell_slot5"`
}
