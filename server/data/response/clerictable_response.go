package response

type ClericTableResponse struct {
	Id           int `json:"id"`
	ProfBonus    int `json:"prof_bonus"`
	CantripKnown int `json:"cantrip_known"`
	SpellSlot1   int `json:"spell_slot1"`
	SpellSlot2   int `json:"spell_slot2"`
	SpellSlot3   int `json:"spell_slot3"`
	SpellSlot4   int `json:"spell_slot4"`
	SpellSlot5   int `json:"spell_slot5"`
	SpellSlot6   int `json:"spell_slot6"`
	SpellSlot7   int `json:"spell_slot7"`
	SpellSlot8   int `json:"spell_slot8"`
	SpellSlot9   int `json:"spell_slot9"`
}
