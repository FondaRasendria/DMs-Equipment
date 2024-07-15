package handler

import (
	"golang-crud/config"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/model"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func MonsterDmgResisHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.MonsterDamageResistences
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func MonsterDmgResisHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.MonsterDmgResisCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(request)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	model := model.MonsterDamageResistences{
		MonsterId: request.MonsterId,
		DamageId:  request.DamageId,
	}

	errCreate := config.DB.Create(&model).Error
	if errCreate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    model,
	})
}

func MonsterDmgResisHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.MonsterDamageResistences
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.MonsterDmgResisResponse{
		Id:        model.Id,
		MonsterId: model.MonsterId,
		DamageId:  model.DamageId,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func MonsterDmgResisHandlerGetByMonsterId(ctx *fiber.Ctx) error {
	monsterId := ctx.Params("monster_id")

	var models []model.MonsterDamageResistences
	err := config.DB.Where("monster_id = ?", monsterId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.MonsterDmgResisResponse
	for _, model := range models {
		responses = append(responses, response.MonsterDmgResisResponse{
			Id:        model.Id,
			MonsterId: model.MonsterId,
			DamageId:  model.DamageId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func MonsterDmgResisHandlerGetByDamageId(ctx *fiber.Ctx) error {
	damageId := ctx.Params("damage_id")

	var models []model.MonsterDamageResistences
	err := config.DB.Where("damage_id = ?", damageId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.MonsterDmgResisResponse
	for _, model := range models {
		responses = append(responses, response.MonsterDmgResisResponse{
			Id:        model.Id,
			MonsterId: model.MonsterId,
			DamageId:  model.DamageId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func MonsterDmgResisHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.MonsterDmgResisUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.MonsterDamageResistences

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.MonsterId != 0 {
		model.MonsterId = request.MonsterId
	}
	if request.DamageId != 0 {
		model.DamageId = request.DamageId
	}

	errUpdate := config.DB.Save(&model).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    model,
	})
}

func MonsterDmgResisHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.MonsterDamageResistences

	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	errDelete := config.DB.Delete(&model).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "data has been deleted",
	})
}
