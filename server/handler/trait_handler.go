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

func TraitHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Trait
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func TraitHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.TraitCreateRequest)
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

	newTrait := model.Trait{
		ParentId:    request.ParentId,
		Name:        request.Name,
		Description: request.Description,
	}

	errCreate := config.DB.Create(&newTrait).Error
	if errCreate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newTrait,
	})
}

func TraitHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Trait
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.TraitResponse{
		Id:          model.Id,
		ParentId:    model.ParentId,
		Name:        model.Name,
		Description: model.Description,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func TraitHandlerGetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var models []model.Trait
	err := config.DB.Where("name = ?", name).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.TraitResponse
	for _, model := range models {
		responses = append(responses, response.TraitResponse{
			Id:          model.Id,
			ParentId:    model.ParentId,
			Name:        model.Name,
			Description: model.Description,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func TraitHandlerGetByParentId(ctx *fiber.Ctx) error {
	parentId := ctx.Params("parent_id")

	var models []model.Trait
	err := config.DB.Where("parent_id = ?", parentId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.TraitResponse
	for _, model := range models {
		responses = append(responses, response.TraitResponse{
			Id:          model.Id,
			ParentId:    model.ParentId,
			Name:        model.Name,
			Description: model.Description,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func TraitHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.TraitUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Trait

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	model.ParentId = request.ParentId
	model.Name = request.Name
	model.Description = request.Description

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

func TraitHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Trait

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
