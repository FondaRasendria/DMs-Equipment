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

func PlayerLanguagesHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.PlayerLanguages
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func PlayerLanguagesHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.PlayerLanguageCreateRequest)
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

	model := model.PlayerLanguages{
		PlayerId:   request.PlayerId,
		LanguageId: request.LanguageId,
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

func PlayerLanguagesHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.PlayerLanguages
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.PlayerLanguageResponse{
		Id:         model.Id,
		PlayerId:   model.PlayerId,
		LanguageId: model.LanguageId,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func PlayerLanguagesHandlerGetByPlayerId(ctx *fiber.Ctx) error {
	playerId := ctx.Params("player_id")

	var models []model.PlayerLanguages
	err := config.DB.Where("player_id = ?", playerId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.PlayerLanguageResponse
	for _, model := range models {
		responses = append(responses, response.PlayerLanguageResponse{
			Id:         model.Id,
			PlayerId:   model.PlayerId,
			LanguageId: model.LanguageId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func PlayerLanguagesHandlerGetByLanguageId(ctx *fiber.Ctx) error {
	languageId := ctx.Params("language_id")

	var models []model.PlayerLanguages
	err := config.DB.Where("language_id = ?", languageId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.PlayerLanguageResponse
	for _, model := range models {
		responses = append(responses, response.PlayerLanguageResponse{
			Id:         model.Id,
			PlayerId:   model.PlayerId,
			LanguageId: model.LanguageId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func PlayerLanguagesHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.PlayerLanguageUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.PlayerLanguages

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.PlayerId != 0 {
		model.PlayerId = request.PlayerId
	}
	if request.LanguageId != 0 {
		model.LanguageId = request.LanguageId
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

func PlayerLanguagesHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.PlayerLanguages

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
