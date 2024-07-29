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

func CampaignHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Campaign
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func CampaignHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.CampaignCreateRequest)
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

	model := model.Campaign{
		Title: request.Title,
		Notes: request.Notes,
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

func CampaignHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Campaign
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.CampaignResponse{
		Id:    model.Id,
		Title: model.Title,
		Notes: model.Notes,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func CampaignHandlerGetByTitle(ctx *fiber.Ctx) error {
	title := ctx.Params("title")

	var models []model.Campaign
	err := config.DB.Where("title = ?", title).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.CampaignResponse
	for _, model := range models {
		responses = append(responses, response.CampaignResponse{
			Id:    model.Id,
			Title: model.Title,
			Notes: model.Notes,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func CampaignHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.CampaignUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Campaign

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.Title != "" {
		model.Title = request.Title
	}
	if request.Notes != "" {
		model.Notes = request.Notes
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

func CampaignHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Campaign

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
