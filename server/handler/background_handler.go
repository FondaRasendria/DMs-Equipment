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

func BackgroundHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Background
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func BackgroundHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.BackgroundCreateRequest)
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

	model := model.Background{
		Name:        request.Name,
		Description: request.Description,
		Source:      request.Source,
		SkillProf:   request.SkillProf,
		ToolProf:    request.ToolProf,
		Language:    request.Language,
		Equipment:   request.Equipment,
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

func BackgroundHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Background
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.BackgroundResponse{
		Id:          model.Id,
		Name:        model.Name,
		Description: model.Description,
		Source:      model.Source,
		SkillProf:   model.SkillProf,
		ToolProf:    model.ToolProf,
		Language:    model.Language,
		Equipment:   model.Equipment,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func BackgroundHandlerGetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var models []model.Background
	err := config.DB.Where("name = ?", name).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.BackgroundResponse
	for _, model := range models {
		responses = append(responses, response.BackgroundResponse{
			Id:          model.Id,
			Name:        model.Name,
			Description: model.Description,
			Source:      model.Source,
			SkillProf:   model.SkillProf,
			ToolProf:    model.ToolProf,
			Language:    model.Language,
			Equipment:   model.Equipment,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func BackgroundHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.BackgroundUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Background

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
	if request.Description != "" {
		model.Description = request.Description
	}
	if request.Source != "" {
		model.Source = request.Source
	}
	if request.SkillProf != "" {
		model.SkillProf = request.SkillProf
	}
	if request.ToolProf != "" {
		model.ToolProf = request.ToolProf
	}
	if request.Language != "" {
		model.Language = request.Language
	}
	if request.Equipment != "" {
		model.Equipment = request.Equipment
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

func BackgroundHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Background

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
