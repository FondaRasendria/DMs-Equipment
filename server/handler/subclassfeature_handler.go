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

func SubclassFeatureHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.SubclassFeature
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func SubclassFeatureHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.SubclassFeatureCreateRequest)
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

	model := model.SubclassFeature{
		ParentId:    request.ParentId,
		Level:       request.Level,
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

func SubclassFeatureHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.SubclassFeature
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.SubclassFeatureResponse{
		Id:          model.Id,
		ParentId:    model.ParentId,
		Level:       model.Level,
		Description: model.Description,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func SubclassFeatureHandlerGetByParentId(ctx *fiber.Ctx) error {
	parentId := ctx.Params("parent_id")

	var models []model.SubclassFeature
	err := config.DB.Where("parent_id = ?", parentId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.SubclassFeatureResponse
	for _, model := range models {
		responses = append(responses, response.SubclassFeatureResponse{
			Id:          model.Id,
			ParentId:    model.ParentId,
			Level:       model.Level,
			Description: model.Description,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func SubclassFeatureHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.SubclassFeatureUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.SubclassFeature

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
	if request.Level != 0 {
		model.Level = request.Level
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

func SubclassFeatureHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.SubclassFeature

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
