package router

import (
	"golang-crud/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world",
		})
	})

	//Actions
	r.Get("/api/action", handler.ActionHandlerGetAll)
	r.Get("/api/action/id/:id", handler.ActionHandlerGetById)
	r.Get("/api/action/parentid/:parent_id", handler.ActionHandlerGetByParentId)
	r.Post("/api/action", handler.ActionHandlerCreate)
	r.Put("/api/action/:id", handler.ActionHandlerUpdate)
	r.Delete("/api/action/:id", handler.ActionHandlerDelete)

	//Monsters
	r.Get("/api/monster", handler.MonsterHandlerGetAll)
	r.Get("/api/monster/id/:id", handler.MonsterHandlerGetById)
	r.Get("/api/monster/name/:name", handler.MonsterHandlerGetByName)
	r.Post("/api/monster", handler.MonsterHandlerCreate)
	r.Put("/api/monster/:id", handler.MonsterHandlerUpdate)
	r.Delete("/api/monster/:id", handler.MonsterHandlerDelete)

	//MonsterTraits
	r.Get("/api/monstertrait", handler.MonsterTraitHandlerGetAll)
	r.Get("/api/monstertrait/id/:id", handler.MonsterTraitHandlerGetById)
	r.Get("/api/monstertrait/parentid/:parent_id", handler.MonsterTraitHandlerGetByParentId)
	r.Get("/api/monstertrait/name/:name", handler.MonsterTraitHandlerGetByName)
	r.Post("/api/monstertrait", handler.MonsterTraitHandlerCreate)
	r.Put("/api/monstertrait/:id", handler.MonsterTraitHandlerUpdate)
	r.Delete("/api/monstertrait/:id", handler.MonsterTraitHandlerDelete)

	//Races
	r.Get("/api/race", handler.RaceHandlerGetAll)
	r.Get("/api/race/id/:id", handler.RaceHandlerGetById)
	r.Get("/api/race/name/:name", handler.RaceHandlerGetByName)
	r.Post("/api/race", handler.RaceHandlerCreate)
	r.Put("/api/race/:id", handler.RaceHandlerUpdate)
	r.Delete("/api/race/:id", handler.RaceHandlerDelete)

	//Subraces
	r.Get("/api/subrace", handler.SubraceHandlerGetAll)
	r.Get("/api/subrace/id/:id", handler.SubraceHandlerGetById)
	r.Get("/api/subrace/parentid/:parent_id", handler.SubraceHandlerGetByParentId)
	r.Post("/api/subrace", handler.SubraceHandlerCreate)
	r.Put("/api/subrace/:id", handler.SubraceHandlerUpdate)
	r.Delete("/api/subrace/:id", handler.SubraceHandlerDelete)

	//Subtraits
	r.Get("/api/subtrait", handler.SubtraitHandlerGetAll)
	r.Get("/api/subtrait/id/:id", handler.SubtraitHandlerGetById)
	r.Get("/api/subtrait/parentid/:parent_id", handler.SubtraitHandlerGetByParentId)
	r.Post("/api/subtrait", handler.SubtraitHandlerCreate)
	r.Put("/api/subtrait/:id", handler.SubtraitHandlerUpdate)
	r.Delete("/api/subtrait/:id", handler.SubtraitHandlerDelete)

	//Traits
	r.Get("/api/trait", handler.TraitHandlerGetAll)
	r.Get("/api/trait/id/:id", handler.TraitHandlerGetById)
	r.Get("/api/trait/parentid/:parent_id", handler.TraitHandlerGetByParentId)
	r.Get("/api/trait/name/:name", handler.TraitHandlerGetByName)
	r.Post("/api/trait", handler.TraitHandlerCreate)
	r.Put("/api/trait/:id", handler.TraitHandlerUpdate)
	r.Delete("/api/trait/:id", handler.TraitHandlerDelete)

	//ConditionTypes
	r.Get("/api/condition", handler.ConditionTypeHandlerGetAll)
	r.Get("/api/condition/id/:id", handler.ConditionTypeHandlerGetById)
	r.Get("/api/condition/name/:name", handler.ConditionTypeHandlerGetByName)
	r.Post("/api/condition", handler.ConditionTypeHandlerCreate)
	r.Put("/api/condition/:id", handler.ConditionTypeHandlerUpdate)
	r.Delete("/api/condition/:id", handler.ConditionTypeHandlerDelete)

	//DamageTypes
	r.Get("/api/damage", handler.DamageTypeHandlerGetAll)
	r.Get("/api/damage/id/:id", handler.DamageTypeHandlerGetById)
	r.Get("/api/damage/name/:name", handler.DamageTypeHandlerGetByName)
	r.Post("/api/damage", handler.DamageTypeHandlerCreate)
	r.Put("/api/damage/:id", handler.DamageTypeHandlerUpdate)
	r.Delete("/api/damage/:id", handler.DamageTypeHandlerDelete)

	//Languages
	r.Get("/api/language", handler.LanguageHandlerGetAll)
	r.Get("/api/language/id/:id", handler.LanguageHandlerGetById)
	r.Get("/api/language/name/:name", handler.LanguageHandlerGetByName)
	r.Post("/api/language", handler.LanguageHandlerCreate)
	r.Put("/api/language/:id", handler.LanguageHandlerUpdate)
	r.Delete("/api/language/:id", handler.LanguageHandlerDelete)

	//SavingThrows
	r.Get("/api/saving", handler.SavingThrowHandlerGetAll)
	r.Get("/api/saving/id/:id", handler.SavingThrowHandlerGetById)
	r.Get("/api/saving/name/:name", handler.SavingThrowHandlerGetByName)
	r.Post("/api/saving", handler.SavingThrowHandlerCreate)
	r.Put("/api/saving/:id", handler.SavingThrowHandlerUpdate)
	r.Delete("/api/saving/:id", handler.SavingThrowHandlerDelete)

	//Skills
	r.Get("/api/skill", handler.SkillTypeHandlerGetAll)
	r.Get("/api/skill/id/:id", handler.SkillTypeHandlerGetById)
	r.Get("/api/skill/name/:name", handler.SkillTypeHandlerGetByName)
	r.Post("/api/skill", handler.SkillTypeHandlerCreate)
	r.Put("/api/skill/:id", handler.SkillTypeHandlerUpdate)
	r.Delete("/api/skill/:id", handler.SkillTypeHandlerDelete)

	//Spells
	r.Get("/api/spell", handler.SpellHandlerGetAll)
	r.Get("/api/spell/id/:id", handler.SpellHandlerGetById)
	r.Get("/api/spell/name/:name", handler.SpellHandlerGetByName)
	r.Post("/api/spell", handler.SpellHandlerCreate)
	r.Put("/api/spell/:id", handler.SpellHandlerUpdate)
	r.Delete("/api/spell/:id", handler.SpellHandlerDelete)

	//Backgrounds
	r.Get("/api/background", handler.BackgroundHandlerGetAll)
	r.Get("/api/background/id/:id", handler.BackgroundHandlerGetById)
	r.Get("/api/background/name/:name", handler.BackgroundHandlerGetByName)
	r.Post("/api/background", handler.BackgroundHandlerCreate)
	r.Put("/api/background/:id", handler.BackgroundHandlerUpdate)
	r.Delete("/api/background/:id", handler.BackgroundHandlerDelete)

	//Classes
	r.Get("/api/class", handler.ClassHandlerGetAll)
	r.Get("/api/class/id/:id", handler.ClassHandlerGetById)
	r.Get("/api/class/name/:name", handler.ClassHandlerGetByName)
	r.Post("/api/class", handler.ClassHandlerCreate)
	r.Put("/api/class/:id", handler.ClassHandlerUpdate)
	r.Delete("/api/class/:id", handler.ClassHandlerDelete)

	//ClassFeatures
	r.Get("/api/classfeature", handler.ClassFeatureHandlerGetAll)
	r.Get("/api/classfeature/id/:id", handler.ClassFeatureHandlerGetById)
	r.Get("/api/classfeature/parentid/:parent_id", handler.ClassFeatureHandlerGetByParentId)
	r.Post("/api/classfeature", handler.ClassFeatureHandlerCreate)
	r.Put("/api/classfeature/:id", handler.ClassFeatureHandlerUpdate)
	r.Delete("/api/classfeature/:id", handler.ClassFeatureHandlerDelete)

	//Subclasses
	r.Get("/api/subclass", handler.SubclassHandlerGetAll)
	r.Get("/api/subclass/id/:id", handler.SubclassHandlerGetById)
	r.Get("/api/subclass/name/:name", handler.SubclassHandlerGetByName)
	r.Get("/api/subclass/parentid/:parent_id", handler.SubclassHandlerGetByParentId)
	r.Post("/api/subclass", handler.SubclassHandlerCreate)
	r.Put("/api/subclass/:id", handler.SubclassHandlerUpdate)
	r.Delete("/api/subclass/:id", handler.SubclassHandlerDelete)

	//SubclassFeatures
	r.Get("/api/subclassfeature", handler.SubclassFeatureHandlerGetAll)
	r.Get("/api/subclassfeature/id/:id", handler.SubclassFeatureHandlerGetById)
	r.Get("/api/subclassfeature/parentid/:parent_id", handler.SubclassFeatureHandlerGetByParentId)
	r.Post("/api/subclassfeature", handler.SubclassFeatureHandlerCreate)
	r.Put("/api/subclassfeature/:id", handler.SubclassFeatureHandlerUpdate)
	r.Delete("/api/subclassfeature/:id", handler.SubclassFeatureHandlerDelete)

	//Campaign
	r.Get("/api/campaign", handler.CampaignHandlerGetAll)
	r.Get("/api/campaign/id/:id", handler.CampaignHandlerGetById)
	r.Get("/api/campaign/title/:title", handler.CampaignHandlerGetByTitle)
	r.Post("/api/campaign", handler.CampaignHandlerCreate)
	r.Put("/api/campaign/:id", handler.CampaignHandlerUpdate)
	r.Delete("/api/campaign/:id", handler.CampaignHandlerDelete)

	//Encounter
	r.Get("/api/encounter", handler.EncounterHandlerGetAll)
	r.Get("/api/encounter/id/:id", handler.EncounterHandlerGetById)
	r.Get("/api/encounter/name/:name", handler.EncounterHandlerGetByName)
	r.Post("/api/encounter", handler.EncounterHandlerCreate)
	r.Put("/api/encounter/:id", handler.EncounterHandlerUpdate)
	r.Delete("/api/encounter/:id", handler.EncounterHandlerDelete)

	//EncounterMonsters
	r.Get("/api/encountermonster", handler.EncounterMonsterHandlerGetAll)
	r.Get("/api/encountermonster/id/:id", handler.EncounterMonsterHandlerGetById)
	r.Get("/api/encountermonster/encounterid/:encounter_id", handler.EncounterMonsterHandlerGetByEncounterId)
	r.Get("/api/encountermonster/monsterid/:monster_id", handler.EncounterMonsterHandlerGetByMonsterId)
	r.Post("/api/encountermonster", handler.EncounterMonsterHandlerCreate)
	r.Put("/api/encountermonster/:id", handler.EncounterMonsterHandlerUpdate)
	r.Delete("/api/encountermonster/:id", handler.EncounterMonsterHandlerDelete)

	//EncounterPlayers
	r.Get("/api/encounterplayer", handler.EncounterPlayerHandlerGetAll)
	r.Get("/api/encounterplayer/id/:id", handler.EncounterPlayerHandlerGetById)
	r.Get("/api/encounterplayer/encounterid/:encounter_id", handler.EncounterPlayerHandlerGetByEncounterId)
	r.Get("/api/encounterplayer/playerid/:player_id", handler.EncounterPlayerHandlerGetByPlayerId)
	r.Post("/api/encounterplayer", handler.EncounterPlayerHandlerCreate)
	r.Put("/api/encounterplayer/:id", handler.EncounterPlayerHandlerUpdate)
	r.Delete("/api/encounterplayer/:id", handler.EncounterPlayerHandlerDelete)

	//MonsterConditionImmunities
	r.Get("/api/monsterconimmun", handler.MonsterConditionImmunitiesHandlerGetAll)
	r.Get("/api/monsterconimmun/id/:id", handler.MonsterConditionImmunitiesHandlerGetById)
	r.Get("/api/monsterconimmun/conditionid/:condition_id", handler.MonsterConditionImmunitiesHandlerGetByConditionId)
	r.Get("/api/monsterconimmun/monsterid/:monster_id", handler.MonsterConditionImmunitiesHandlerGetByMonsterId)
	r.Post("/api/monsterconimmun", handler.MonsterConditionImmunitiesHandlerCreate)
	r.Put("/api/monsterconimmun/:id", handler.MonsterConditionImmunitiesHandlerUpdate)
	r.Delete("/api/monsterconimmun/:id", handler.MonsterConditionImmunitiesHandlerDelete)

	//MonsterDamageImmunities
	r.Get("/api/monsterdmgimmun", handler.MonsterDmgImmunHandlerGetAll)
	r.Get("/api/monsterdmgimmun/id/:id", handler.MonsterDmgImmunHandlerGetById)
	r.Get("/api/monsterdmgimmun/damageid/:damage_id", handler.MonsterDmgImmunHandlerGetByDamageId)
	r.Get("/api/monsterdmgimmun/monsterid/:monster_id", handler.MonsterDmgImmunHandlerGetByMonsterId)
	r.Post("/api/monsterdmgimmun", handler.MonsterDmgImmunHandlerCreate)
	r.Put("/api/monsterdmgimmun/:id", handler.MonsterDmgImmunHandlerUpdate)
	r.Delete("/api/monsterdmgimmun/:id", handler.MonsterDmgImmunHandlerDelete)

	//MonsterDamageResistences
	r.Get("/api/monsterdmgresis", handler.MonsterDmgResisHandlerGetAll)
	r.Get("/api/monsterdmgresis/id/:id", handler.MonsterDmgResisHandlerGetById)
	r.Get("/api/monsterdmgresis/damageid/:damage_id", handler.MonsterDmgResisHandlerGetByDamageId)
	r.Get("/api/monsterdmgresis/monsterid/:monster_id", handler.MonsterDmgResisHandlerGetByMonsterId)
	r.Post("/api/monsterdmgresis", handler.MonsterDmgResisHandlerCreate)
	r.Put("/api/monsterdmgresis/:id", handler.MonsterDmgResisHandlerUpdate)
	r.Delete("/api/monsterdmgresis/:id", handler.MonsterDmgResisHandlerDelete)

	//MonsterDamageVulnerabilities
	r.Get("/api/monsterdmgvulne", handler.MonsterDmgVulneHandlerGetAll)
	r.Get("/api/monsterdmgvulne/id/:id", handler.MonsterDmgVulneHandlerGetById)
	r.Get("/api/monsterdmgvulne/damageid/:damage_id", handler.MonsterDmgVulneHandlerGetByDamageId)
	r.Get("/api/monsterdmgvulne/monsterid/:monster_id", handler.MonsterDmgVulneHandlerGetByMonsterId)
	r.Post("/api/monsterdmgvulne", handler.MonsterDmgVulneHandlerCreate)
	r.Put("/api/monsterdmgvulne/:id", handler.MonsterDmgVulneHandlerUpdate)
	r.Delete("/api/monsterdmgvulne/:id", handler.MonsterDmgVulneHandlerDelete)

	//MonsterLanguages
	r.Get("/api/monsterlanguage", handler.MonsterLanguagessHandlerGetAll)
	r.Get("/api/monsterlanguage/id/:id", handler.MonsterLanguagesHandlerGetById)
	r.Get("/api/monsterlanguage/languageid/:language_id", handler.MonsterLanguagessHandlerGetByLanguagessId)
	r.Get("/api/monsterlanguage/monsterid/:monster_id", handler.MonsterLanguagessHandlerGetByLanguagessId)
	r.Post("/api/monsterlanguage", handler.MonsterLanguagesHandlerCreate)
	r.Put("/api/monsterlanguage/:id", handler.MonsterLanguagessHandlerUpdate)
	r.Delete("/api/monsterlanguage/:id", handler.MonsterLanguagessHandlerDelete)

	//MonsterSavings
	r.Get("/api/monstersaving", handler.MonsterSavingProfHandlerGetAll)
	r.Get("/api/monstersaving/id/:id", handler.MonsterSavingProfHandlerGetById)
	r.Get("/api/monstersaving/savingid/:saving_id", handler.MonsterSavingProfHandlerGetBySavingId)
	r.Get("/api/monstersaving/monsterid/:monster_id", handler.MonsterSavingProfHandlerGetByMonsterId)
	r.Post("/api/monstersaving", handler.MonsterSavingProfHandlerCreate)
	r.Put("/api/monstersaving/:id", handler.MonsterSavingProfHandlerUpdate)
	r.Delete("/api/monstersaving/:id", handler.MonsterSavingProfHandlerDelete)

	//MonsterSkills
	r.Get("/api/monsterskill", handler.MonsterSkillProfHandlerGetAll)
	r.Get("/api/monsterskill/id/:id", handler.MonsterSkillProfHandlerGetById)
	r.Get("/api/monsterskill/skillid/:skill_id", handler.MonsterSkillProfHandlerGetBySkillId)
	r.Get("/api/monsterskill/monsterid/:monster_id", handler.MonsterSkillProfHandlerGetByMonsterId)
	r.Post("/api/monsterskill", handler.MonsterSkillProfHandlerCreate)
	r.Put("/api/monsterskill/:id", handler.MonsterSkillProfHandlerUpdate)
	r.Delete("/api/monsterskill/:id", handler.MonsterSkillProfHandlerDelete)

	//Player
	r.Get("/api/player", handler.PlayerHandlerGetAll)
	r.Get("/api/player/id/:id", handler.PlayerHandlerGetById)
	r.Get("/api/player/name/:name", handler.PlayerHandlerGetByName)
	r.Post("/api/player", handler.PlayerHandlerCreate)
	r.Put("/api/player/:id", handler.PlayerHandlerUpdate)
	r.Delete("/api/player/:id", handler.PlayerHandlerDelete)

	//PlayerConditionImmunities
	r.Get("/api/playerconimmun", handler.PlayerConditionImmunHandlerGetAll)
	r.Get("/api/playerconimmun/id/:id", handler.PlayerConditionImmunHandlerGetById)
	r.Get("/api/playerconimmun/conditionid/:condition_id", handler.PlayerConditionImmunHandlerGetByConditionId)
	r.Get("/api/playerconimmun/playerid/:player_id", handler.PlayerConditionImmunHandlerGetByPlayerId)
	r.Post("/api/playerconimmun", handler.PlayerConditionImmunHandlerCreate)
	r.Put("/api/playerconimmun/:id", handler.PlayerConditionImmunHandlerUpdate)
	r.Delete("/api/playerconimmun/:id", handler.PlayerConditionImmunHandlerDelete)

	//PlayerDamageImmunities
	r.Get("/api/playerdmgimmun", handler.PlayerDamageImmunHandlerGetAll)
	r.Get("/api/playerdmgimmun/id/:id", handler.PlayerDamageImmunHandlerGetById)
	r.Get("/api/playerdmgimmun/damageid/:damage_id", handler.PlayerDamageImmunHandlerGetByDamageId)
	r.Get("/api/playerdmgimmun/playerid/:player_id", handler.PlayerDamageImmunHandlerGetByPlayerId)
	r.Post("/api/playerdmgimmun", handler.PlayerDamageImmunHandlerCreate)
	r.Put("/api/playerdmgimmun/:id", handler.PlayerDamageImmunHandlerUpdate)
	r.Delete("/api/playerdmgimmun/:id", handler.PlayerDamageImmunHandlerDelete)

	//PlayerDamageResistences
	r.Get("/api/playerdmgresis", handler.PlayerDamageResisHandlerGetAll)
	r.Get("/api/playerdmgresis/id/:id", handler.PlayerDamageResisHandlerGetById)
	r.Get("/api/playerdmgresis/damageid/:damage_id", handler.PlayerDamageResisHandlerGetByDamageId)
	r.Get("/api/playerdmgresis/playerid/:player_id", handler.PlayerDamageResisHandlerGetByPlayerId)
	r.Post("/api/playerdmgresis", handler.PlayerDamageResisHandlerCreate)
	r.Put("/api/playerdmgresis/:id", handler.PlayerDamageResisHandlerUpdate)
	r.Delete("/api/playerdmgresis/:id", handler.PlayerDamageResisHandlerDelete)

	//PlayerDamageVulnerabilities
	r.Get("/api/playerdmgvulne", handler.PlayerDamageResisHandlerGetAll)
	r.Get("/api/playerdmgvulne/id/:id", handler.PlayerDamageResisHandlerGetById)
	r.Get("/api/playerdmgvulne/damageid/:damage_id", handler.PlayerDamageResisHandlerGetByDamageId)
	r.Get("/api/playerdmgvulne/playerid/:player_id", handler.PlayerDamageResisHandlerGetByPlayerId)
	r.Post("/api/playerdmgvulne", handler.PlayerDamageResisHandlerCreate)
	r.Put("/api/playerdmgvulne/:id", handler.PlayerDamageResisHandlerUpdate)
	r.Delete("/api/playerdmgvulne/:id", handler.PlayerDamageVulneHandlerDelete)

	//PlayerLanguages
	r.Get("/api/playerlanguage", handler.PlayerLanguagesHandlerGetAll)
	r.Get("/api/playerlanguage/id/:id", handler.PlayerLanguagesHandlerGetById)
	r.Get("/api/playerlanguage/languageid/:language_id", handler.PlayerLanguagesHandlerGetByLanguageId)
	r.Get("/api/playerlanguage/playerid/:player_id", handler.PlayerLanguagesHandlerGetByPlayerId)
	r.Post("/api/playerlanguage", handler.PlayerLanguagesHandlerCreate)
	r.Put("/api/playerlanguage/:id", handler.PlayerLanguagesHandlerUpdate)
	r.Delete("/api/playerlanguage/:id", handler.PlayerLanguagesHandlerDelete)

	//PlayerSavings
	r.Get("/api/playersaving", handler.PlayerSavingProfHandlerGetAll)
	r.Get("/api/playersaving/id/:id", handler.PlayerSavingProfHandlerGetById)
	r.Get("/api/playersaving/savingid/:saving_id", handler.PlayerSavingProfHandlerGetBySavingId)
	r.Get("/api/playersaving/playerid/:player_id", handler.PlayerSavingProfHandlerGetByPlayerId)
	r.Post("/api/playersaving", handler.PlayerSavingProfHandlerCreate)
	r.Put("/api/playersaving/:id", handler.PlayerSavingProfHandlerUpdate)
	r.Delete("/api/playersaving/:id", handler.PlayerSavingProfHandlerDelete)

	//PlayerSkills
	r.Get("/api/playerskill", handler.PlayerSkillHandlerGetAll)
	r.Get("/api/playerskill/id/:id", handler.PlayerSkillHandlerGetById)
	r.Get("/api/playerskill/skillid/:skill_id", handler.PlayerSkillHandlerGetBySkillId)
	r.Get("/api/playerskill/playerid/:player_id", handler.PlayerSkillHandlerGetByPlayerId)
	r.Post("/api/playerskill", handler.PlayerSkillHandlerCreate)
	r.Put("/api/playerskill/:id", handler.PlayerSkillHandlerUpdate)
	r.Delete("/api/playerskill/:id", handler.PlayerSkillHandlerDelete)

	//ArtificerTable
	r.Get("/api/artificertable", handler.ArtificerTableHandlerGetAll)
	r.Get("/api/artificertable/id/:id", handler.ArtificerTableHandlerGetById)
	r.Post("/api/artificertable", handler.ArtificerTableHandlerCreate)
	r.Put("/api/artificertable/:id", handler.ArtificerTableHandlerUpdate)
	r.Delete("/api/artificertable/:id", handler.ArtificerTableHandlerDelete)

	//BarbarianTable
	r.Get("/api/barbariantable", handler.BarbarianTableHandlerGetAll)
	r.Get("/api/barbariantable/id/:id", handler.BarbarianTableHandlerGetById)
	r.Post("/api/barbariantable", handler.BarbarianTableHandlerCreate)
	r.Put("/api/barbariantable/:id", handler.BarbarianTableHandlerUpdate)
	r.Delete("/api/barbariantable/:id", handler.BarbarianTableHandlerDelete)

	//BardTable
	r.Get("/api/bardtable", handler.BardTableHandlerGetAll)
	r.Get("/api/bardtable/id/:id", handler.BardTableHandlerGetById)
	r.Post("/api/bardtable", handler.BardTableHandlerCreate)
	r.Put("/api/bardtable/:id", handler.BardTableHandlerUpdate)
	r.Delete("/api/bardtable/:id", handler.BardTableHandlerDelete)

	//BloodHunterTable
	r.Get("/api/bloodhuntertable", handler.BloodHunterTableHandlerGetAll)
	r.Get("/api/bloodhuntertable/id/:id", handler.BloodHunterTableHandlerGetById)
	r.Post("/api/bloodhuntertable", handler.BloodHunterTableHandlerCreate)
	r.Put("/api/bloodhuntertable/:id", handler.BloodHunterTableHandlerUpdate)
	r.Delete("/api/bloodhuntertable/:id", handler.BloodHunterTableHandlerDelete)

	//ClericTable
	r.Get("/api/clerictable", handler.ClericTableHandlerGetAll)
	r.Get("/api/clerictable/id/:id", handler.ClericTableHandlerGetById)
	r.Post("/api/clerictable", handler.ClericTableHandlerCreate)
	r.Put("/api/clerictable/:id", handler.ClericTableHandlerUpdate)
	r.Delete("/api/clerictable/:id", handler.ClericTableHandlerDelete)

	//DruidTable
	r.Get("/api/druidtable", handler.DruidTableHandlerGetAll)
	r.Get("/api/druidtable/id/:id", handler.DruidTableHandlerGetById)
	r.Post("/api/druidtable", handler.DruidTableHandlerCreate)
	r.Put("/api/druidtable/:id", handler.DruidTableHandlerUpdate)
	r.Delete("/api/druidtable/:id", handler.DruidTableHandlerDelete)

	//FighterTable
	r.Get("/api/fightertable", handler.FighterTableHandlerGetAll)
	r.Get("/api/fightertable/id/:id", handler.FighterTableHandlerGetById)
	r.Post("/api/fightertable", handler.FighterTableHandlerCreate)
	r.Put("/api/fightertable/:id", handler.FighterTableHandlerUpdate)
	r.Delete("/api/fightertable/:id", handler.FighterTableHandlerDelete)

	//MonkTable
	r.Get("/api/monktable", handler.MonkTableHandlerGetAll)
	r.Get("/api/monktable/id/:id", handler.MonkTableHandlerGetById)
	r.Post("/api/monktable", handler.MonkTableHandlerCreate)
	r.Put("/api/monktable/:id", handler.MonkTableHandlerUpdate)
	r.Delete("/api/monktable/:id", handler.MonkTableHandlerDelete)

	//PaladinTable
	r.Get("/api/paladintable", handler.PaladinTableHandlerGetAll)
	r.Get("/api/paladintable/id/:id", handler.PaladinTableHandlerGetById)
	r.Post("/api/paladintable", handler.PaladinTableHandlerCreate)
	r.Put("/api/paladintable/:id", handler.PaladinTableHandlerUpdate)
	r.Delete("/api/paladintable/:id", handler.PaladinTableHandlerDelete)

	//RangerTable
	r.Get("/api/rangertable", handler.RangerTableHandlerGetAll)
	r.Get("/api/rangertable/id/:id", handler.RangerTableHandlerGetById)
	r.Post("/api/rangertable", handler.RangerTableHandlerCreate)
	r.Put("/api/rangertable/:id", handler.RangerTableHandlerUpdate)
	r.Delete("/api/rangertable/:id", handler.RangerTableHandlerDelete)

	//RogueTable
	r.Get("/api/roguetable", handler.RogueTableHandlerGetAll)
	r.Get("/api/roguetable/id/:id", handler.RogueTableHandlerGetById)
	r.Post("/api/roguetable", handler.RogueTableHandlerCreate)
	r.Put("/api/roguetable/:id", handler.RogueTableHandlerUpdate)
	r.Delete("/api/roguetable/:id", handler.RogueTableHandlerDelete)

	//SorcererTable
	r.Get("/api/sorcerertable", handler.SorcererTableHandlerGetAll)
	r.Get("/api/sorcerertable/id/:id", handler.SorcererTableHandlerGetById)
	r.Post("/api/sorcerertable", handler.SorcererTableHandlerCreate)
	r.Put("/api/sorcerertable/:id", handler.SorcererTableHandlerUpdate)
	r.Delete("/api/sorcerertable/:id", handler.SorcererTableHandlerDelete)

	//WarlockTable
	r.Get("/api/warlocktable", handler.WarlockTableHandlerGetAll)
	r.Get("/api/warlocktable/id/:id", handler.WarlockTableHandlerGetById)
	r.Post("/api/warlocktable", handler.WarlockTableHandlerCreate)
	r.Put("/api/warlocktable/:id", handler.WarlockTableHandlerUpdate)
	r.Delete("/api/warlocktable/:id", handler.WarlockTableHandlerDelete)

	//WizardTable
	r.Get("/api/wizardtable", handler.WizardTableHandlerGetAll)
	r.Get("/api/wizardtable/id/:id", handler.WizardTableHandlerGetById)
	r.Post("/api/wizardtable", handler.WizardTableHandlerCreate)
	r.Put("/api/wizardtable/:id", handler.WizardTableHandlerUpdate)
	r.Delete("/api/wizardtable/:id", handler.WizardTableHandlerDelete)
}
