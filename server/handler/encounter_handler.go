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

func EncounterHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Encounter
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func EncounterHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.EncounterCreateRequest)
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

	model := model.Encounter{
		CampaignId: request.CampaignId,
		Name:       request.Name,
		Turn:       request.Turn,
		Round:      request.Round,
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

func EncounterHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Encounter
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.EncounterResponse{
		Id:         model.Id,
		CampaignId: model.CampaignId,
		Name:       model.Name,
		Turn:       model.Turn,
		Round:      model.Round,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func EncounterHandlerGetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var models []model.Encounter
	err := config.DB.Where("name = ?", name).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.EncounterResponse
	for _, model := range models {
		responses = append(responses, response.EncounterResponse{
			Id:         model.Id,
			CampaignId: model.CampaignId,
			Name:       model.Name,
			Turn:       model.Turn,
			Round:      model.Round,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func EncounterHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.EncounterUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Encounter

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.CampaignId != 0 {
		model.CampaignId = request.CampaignId
	}
	if request.Name != "" {
		model.Name = request.Name
	}
	if request.Turn != 0 {
		model.Turn = request.Turn
	}
	if request.Round != 0 {
		model.Round = request.Round
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

func EncounterHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Encounter

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
