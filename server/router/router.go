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
	r.Get("/api/action", handler.ActionHandlerGetAll)
	r.Get("/api/action/id/:id", handler.ActionHandlerGetById)
	r.Get("/api/action/parentid/:parent_id", handler.ActionHandlerGetByParentId)
	r.Post("/api/action", handler.ActionHandlerCreate)
	r.Put("/api/action/:id", handler.ActionHandlerUpdate)
	r.Delete("/api/action/:id", handler.ActionHandlerDelete)

	r.Get("/api/monster", handler.MonsterHandlerGetAll)
	r.Get("/api/monster/id/:id", handler.MonsterHandlerGetById)
	r.Get("/api/monster/name/:name", handler.MonsterHandlerGetByName)
	r.Post("/api/monster", handler.MonsterHandlerCreate)
	r.Put("/api/monster/:id", handler.MonsterHandlerUpdate)
	r.Delete("/api/monster/:id", handler.MonsterHandlerDelete)

	r.Get("/api/monstertrait", handler.MonsterTraitHandlerGetAll)
	r.Get("/api/monstertrait/id/:id", handler.MonsterTraitHandlerGetById)
	r.Get("/api/monstertrait/parentid/:parent_id", handler.MonsterTraitHandlerGetByParentId)
	r.Get("/api/monstertrait/name/:name", handler.MonsterTraitHandlerGetByName)
	r.Post("/api/monstertrait", handler.MonsterTraitHandlerCreate)
	r.Put("/api/monstertrait/:id", handler.MonsterTraitHandlerUpdate)
	r.Delete("/api/monstertrait/:id", handler.MonsterTraitHandlerDelete)

	r.Get("/api/race", handler.RaceHandlerGetAll)
	r.Get("/api/race/id/:id", handler.RaceHandlerGetById)
	r.Get("/api/race/name/:name", handler.RaceHandlerGetByName)
	r.Post("/api/race", handler.RaceHandlerCreate)
	r.Put("/api/race/:id", handler.RaceHandlerUpdate)
	r.Delete("/api/race/:id", handler.RaceHandlerDelete)

	r.Get("/api/subrace", handler.SubraceHandlerGetAll)
	r.Get("/api/subrace/id/:id", handler.SubraceHandlerGetById)
	r.Get("/api/subrace/parentid/:parent_id", handler.SubraceHandlerGetByParentId)
	r.Post("/api/subrace", handler.SubraceHandlerCreate)
	r.Put("/api/subrace/:id", handler.SubraceHandlerUpdate)
	r.Delete("/api/subrace/:id", handler.SubraceHandlerDelete)

	r.Get("/api/subtrait", handler.SubtraitHandlerGetAll)
	r.Get("/api/subtrait/id/:id", handler.SubtraitHandlerGetById)
	r.Get("/api/subtrait/parentid/:parent_id", handler.SubtraitHandlerGetByParentId)
	r.Post("/api/subtrait", handler.SubtraitHandlerCreate)
	r.Put("/api/subtrait/:id", handler.SubtraitHandlerUpdate)
	r.Delete("/api/subtrait/:id", handler.SubtraitHandlerDelete)

	r.Get("/api/trait", handler.TraitHandlerGetAll)
	r.Get("/api/trait/id/:id", handler.TraitHandlerGetById)
	r.Get("/api/trait/parentid/:parent_id", handler.TraitHandlerGetByParentId)
	r.Get("/api/trait/name/:name", handler.TraitHandlerGetByName)
	r.Post("/api/trait", handler.TraitHandlerCreate)
	r.Put("/api/trait/:id", handler.TraitHandlerUpdate)
	r.Delete("/api/trait/:id", handler.TraitHandlerDelete)

	r.Get("/api/condition", handler.ConditionTypeHandlerGetAll)
	r.Get("/api/condition/id/:id", handler.ConditionTypeHandlerGetById)
	r.Get("/api/condition/name/:name", handler.ConditionTypeHandlerGetByName)
	r.Post("/api/condition", handler.ConditionTypeHandlerCreate)
	r.Put("/api/condition/:id", handler.ConditionTypeHandlerUpdate)
	r.Delete("/api/condition/:id", handler.ConditionTypeHandlerDelete)

	r.Get("/api/damage", handler.DamageTypeHandlerGetAll)
	r.Get("/api/damage/id/:id", handler.DamageTypeHandlerGetById)
	r.Get("/api/damage/name/:name", handler.DamageTypeHandlerGetByName)
	r.Post("/api/damage", handler.DamageTypeHandlerCreate)
	r.Put("/api/damage/:id", handler.DamageTypeHandlerUpdate)
	r.Delete("/api/damage/:id", handler.DamageTypeHandlerDelete)

	r.Get("/api/language", handler.LanguageHandlerGetAll)
	r.Get("/api/language/id/:id", handler.LanguageHandlerGetById)
	r.Get("/api/language/name/:name", handler.LanguageHandlerGetByName)
	r.Post("/api/language", handler.LanguageHandlerCreate)
	r.Put("/api/language/:id", handler.LanguageHandlerUpdate)
	r.Delete("/api/language/:id", handler.LanguageHandlerDelete)

	r.Get("/api/saving", handler.SavingThrowHandlerGetAll)
	r.Get("/api/saving/id/:id", handler.SavingThrowHandlerGetById)
	r.Get("/api/saving/name/:name", handler.SavingThrowHandlerGetByName)
	r.Post("/api/saving", handler.SavingThrowHandlerCreate)
	r.Put("/api/saving/:id", handler.SavingThrowHandlerUpdate)
	r.Delete("/api/saving/:id", handler.SavingThrowHandlerDelete)

	r.Get("/api/skill", handler.SkillTypeHandlerGetAll)
	r.Get("/api/skill/id/:id", handler.SkillTypeHandlerGetById)
	r.Get("/api/skill/name/:name", handler.SkillTypeHandlerGetByName)
	r.Post("/api/skill", handler.SkillTypeHandlerCreate)
	r.Put("/api/skill/:id", handler.SkillTypeHandlerUpdate)
	r.Delete("/api/skill/:id", handler.SkillTypeHandlerDelete)
}
