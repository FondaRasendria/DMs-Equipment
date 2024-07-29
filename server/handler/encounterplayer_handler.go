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

func EncounterPlayerHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.EncounterPlayer
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func EncounterPlayerHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.EncounterPlayerCreateRequest)
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

	model := model.EncounterPlayer{
		EncounterId: request.EncounterId,
		PlayerId:    request.PlayerId,
		Initiative:  request.Initiative,
		CurrentHP:   request.CurrentHP,
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

func EncounterPlayerHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.EncounterPlayer
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.EncounterPlayerResponse{
		Id:          model.Id,
		EncounterId: model.EncounterId,
		PlayerId:    model.PlayerId,
		Initiative:  model.Initiative,
		CurrentHP:   model.CurrentHP,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func EncounterPlayerHandlerGetByEncounterId(ctx *fiber.Ctx) error {
	encounterId := ctx.Params("encounter_id")

	var models []model.EncounterPlayer
	err := config.DB.Where("encounter_id = ?", encounterId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.EncounterPlayerResponse
	for _, model := range models {
		responses = append(responses, response.EncounterPlayerResponse{
			Id:          model.Id,
			EncounterId: model.EncounterId,
			PlayerId:    model.PlayerId,
			Initiative:  model.Initiative,
			CurrentHP:   model.CurrentHP,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func EncounterPlayerHandlerGetByPlayerId(ctx *fiber.Ctx) error {
	monsterId := ctx.Params("monster_id")

	var models []model.EncounterPlayer
	err := config.DB.Where("monster_id = ?", monsterId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.EncounterPlayerResponse
	for _, model := range models {
		responses = append(responses, response.EncounterPlayerResponse{
			Id:          model.Id,
			EncounterId: model.EncounterId,
			PlayerId:    model.PlayerId,
			Initiative:  model.Initiative,
			CurrentHP:   model.CurrentHP,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func EncounterPlayerHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.EncounterPlayerUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.EncounterPlayer

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.EncounterId != 0 {
		model.EncounterId = request.EncounterId
	}
	if request.PlayerId != 0 {
		model.PlayerId = request.PlayerId
	}
	if request.Initiative != 0 {
		model.Initiative = request.Initiative
	}
	if request.CurrentHP != 0 {
		model.CurrentHP = request.CurrentHP
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

func EncounterPlayerHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.EncounterPlayer

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
