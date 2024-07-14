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

func SubraceHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Subrace
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func SubraceHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.SubraceCreateRequest)
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

	model := model.Subrace{
		ParentId:    request.ParentId,
		Description: request.Description,
		Source:      request.Source,
		Ability:     request.Ability,
		Age:         request.Age,
		Alignment:   request.Alignment,
		Size:        request.Size,
		Speed:       request.Speed,
		Language:    request.Language,
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

func SubraceHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Subrace
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.SubraceResponse{
		Id:          model.Id,
		ParentId:    model.ParentId,
		Description: model.Description,
		Source:      model.Source,
		Ability:     model.Ability,
		Age:         model.Age,
		Alignment:   model.Alignment,
		Size:        model.Size,
		Speed:       model.Speed,
		Language:    model.Language,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func SubraceHandlerGetByParentId(ctx *fiber.Ctx) error {
	parentId := ctx.Params("parent_id")

	var models []model.Subrace
	err := config.DB.Where("parent_id = ?", parentId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.SubraceResponse
	for _, model := range models {
		responses = append(responses, response.SubraceResponse{
			Id:          model.Id,
			ParentId:    model.ParentId,
			Description: model.Description,
			Source:      model.Source,
			Ability:     model.Ability,
			Age:         model.Age,
			Alignment:   model.Alignment,
			Size:        model.Size,
			Speed:       model.Speed,
			Language:    model.Language,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func SubraceHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.SubraceUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Subrace

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	model.ParentId = request.ParentId
	model.Description = request.Description
	model.Source = request.Source
	model.Ability = request.Ability
	model.Age = request.Age
	model.Alignment = request.Alignment
	model.Size = request.Size
	model.Speed = request.Speed
	model.Language = request.Language

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

func SubraceHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Subrace

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
