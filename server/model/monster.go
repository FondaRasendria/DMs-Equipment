package model

type Monster struct {
	Id                    int
	Name                  string
	Type                  string
	Alignment             string
	AC                    int
	FixedHP               int
	DiceHP                string
	Speed                 string
	STR                   int
	DEX                   int
	CON                   int
	INT                   int
	WIS                   int
	CHA                   int
	SavingThrows          string
	Skills                string
	DamageImmunities      string
	DamageVulnerabilities string
	DamageResistences     string
	ConditionImmunities   string
	Senses                string
	Language              string
	CR                    string
	Description           string
	Source                string
}
