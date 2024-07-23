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

func SubclassHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Subclass
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func SubclassHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.SubclassCreateRequest)
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

	model := model.Subclass{
		ParentId:    request.ParentId,
		Name:        request.Name,
		Source:      request.Source,
		Description: request.Description,
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

func SubclassHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Subclass
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.SubclassResponse{
		Id:          model.Id,
		ParentId:    model.ParentId,
		Name:        model.Name,
		Source:      model.Source,
		Description: model.Description,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func SubclassHandlerGetByParentId(ctx *fiber.Ctx) error {
	parentId := ctx.Params("parent_id")

	var models []model.Subclass
	err := config.DB.Where("parent_id = ?", parentId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.SubclassResponse
	for _, model := range models {
		responses = append(responses, response.SubclassResponse{
			Id:          model.Id,
			ParentId:    model.ParentId,
			Name:        model.Name,
			Source:      model.Source,
			Description: model.Description,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func SubclassHandlerGetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var models []model.Subclass
	err := config.DB.Where("name = ?", name).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.SubclassResponse
	for _, model := range models {
		responses = append(responses, response.SubclassResponse{
			Id:          model.Id,
			ParentId:    model.ParentId,
			Name:        model.Name,
			Source:      model.Source,
			Description: model.Description,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func SubclassHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.SubclassUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Subclass

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.ParentId != 0 {
		model.ParentId = request.ParentId
	}
	if request.Name != "" {
		model.Name = request.Name
	}
	if request.Source != "" {
		model.Source = request.Source
	}
	if request.Description != "" {
		model.Description = request.Description
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

func SubclassHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Subclass

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
