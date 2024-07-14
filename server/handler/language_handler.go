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

func LanguageHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Language
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func LanguageHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.LanguageCreateRequest)
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

	model := model.Language{
		Name:        request.Name,
		Alphabet:    request.Alphabet,
		MainSpeaker: request.MainSpeaker,
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

func LanguageHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Language
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.LanguageResponse{
		Id:          model.Id,
		Name:        model.Name,
		Alphabet:    model.Alphabet,
		MainSpeaker: model.MainSpeaker,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func LanguageHandlerGetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var models []model.Language
	err := config.DB.Where("name = ?", name).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.LanguageResponse
	for _, model := range models {
		responses = append(responses, response.LanguageResponse{
			Id:          model.Id,
			Name:        model.Name,
			Alphabet:    model.Alphabet,
			MainSpeaker: model.MainSpeaker,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func LanguageHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.LanguageUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Language

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.Name != "" {
		model.Name = request.Name
	}
	if request.Alphabet != "" {
		model.Alphabet = request.Alphabet
	}
	if request.MainSpeaker != "" {
		model.MainSpeaker = request.MainSpeaker
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

func LanguageHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Language

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
