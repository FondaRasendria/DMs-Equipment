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

func MonsterLanguagessHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.MonsterLanguages
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func MonsterLanguagesHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.MonsterLanguagesCreateRequest)
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

	model := model.MonsterLanguages{
		MonsterId:  request.MonsterId,
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

func MonsterLanguagesHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.MonsterLanguages
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.MonsterLanguagesResponse{
		Id:         model.Id,
		MonsterId:  model.MonsterId,
		LanguageId: model.LanguageId,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func MonsterLanguagesHandlerGetByMonsterId(ctx *fiber.Ctx) error {
	monsterId := ctx.Params("monster_id")

	var models []model.MonsterLanguages
	err := config.DB.Where("monster_id = ?", monsterId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.MonsterLanguagesResponse
	for _, model := range models {
		responses = append(responses, response.MonsterLanguagesResponse{
			Id:         model.Id,
			MonsterId:  model.MonsterId,
			LanguageId: model.LanguageId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func MonsterLanguagessHandlerGetByLanguagessId(ctx *fiber.Ctx) error {
	languageId := ctx.Params("language_id")

	var models []model.MonsterLanguages
	err := config.DB.Where("language_id = ?", languageId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.MonsterLanguagesResponse
	for _, model := range models {
		responses = append(responses, response.MonsterLanguagesResponse{
			Id:         model.Id,
			MonsterId:  model.MonsterId,
			LanguageId: model.LanguageId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func MonsterLanguagessHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.MonsterLanguagesUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.MonsterLanguages

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

func MonsterLanguagessHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.MonsterLanguages

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
