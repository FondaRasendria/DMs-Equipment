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

func SpellHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Spell
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func SpellHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.SpellCreateRequest)
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

	model := model.Spell{
		Name:        request.Name,
		Source:      request.Source,
		School:      request.School,
		Level:       request.Level,
		CastingTime: request.CastingTime,
		Range:       request.Range,
		Components:  request.Components,
		Duration:    request.Duration,
		Description: request.Description,
		HigherLevel: request.HigherLevel,
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

func SpellHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Spell
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.SpellResponse{
		Id:          model.Id,
		Name:        model.Name,
		Source:      model.Source,
		School:      model.School,
		Level:       model.Level,
		CastingTime: model.CastingTime,
		Range:       model.Range,
		Components:  model.Components,
		Duration:    model.Duration,
		Description: model.Description,
		HigherLevel: model.HigherLevel,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func SpellHandlerGetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var models []model.Spell
	err := config.DB.Where("name = ?", name).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.SpellResponse
	for _, model := range models {
		responses = append(responses, response.SpellResponse{
			Id:          model.Id,
			Name:        model.Name,
			Source:      model.Source,
			School:      model.School,
			Level:       model.Level,
			CastingTime: model.CastingTime,
			Range:       model.Range,
			Components:  model.Components,
			Duration:    model.Duration,
			Description: model.Description,
			HigherLevel: model.HigherLevel,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func SpellHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.SpellUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Spell

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
	if request.Source != "" {
		model.Source = request.Source
	}
	if request.School != "" {
		model.School = request.School
	}
	if request.Level != 0 {
		model.Level = request.Level
	}
	if request.CastingTime != "" {
		model.CastingTime = request.CastingTime
	}
	if request.Range != "" {
		model.Range = request.Range
	}
	if request.Components != "" {
		model.Components = request.Components
	}
	if request.Duration != "" {
		model.Duration = request.Duration
	}
	if request.Description != "" {
		model.Description = request.Description
	}
	if request.HigherLevel != "" {
		model.HigherLevel = request.HigherLevel
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

func SpellHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Spell

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
