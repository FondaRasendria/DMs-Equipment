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

func RangerTableHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.RangerTable
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func RangerTableHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.RangerTableCreateRequest)
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

	model := model.RangerTable{
		ProfBonus:  request.ProfBonus,
		SpellKnown: request.SpellKnown,
		SpellSlot1: request.SpellSlot1,
		SpellSlot2: request.SpellSlot2,
		SpellSlot3: request.SpellSlot3,
		SpellSlot4: request.SpellSlot4,
		SpellSlot5: request.SpellSlot5,
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

func RangerTableHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.RangerTable
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.RangerTableResponse{
		Id:         model.Id,
		ProfBonus:  model.ProfBonus,
		SpellKnown: model.SpellKnown,
		SpellSlot1: model.SpellSlot1,
		SpellSlot2: model.SpellSlot2,
		SpellSlot3: model.SpellSlot3,
		SpellSlot4: model.SpellSlot4,
		SpellSlot5: model.SpellSlot5,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func RangerTableHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.RangerTableUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.RangerTable

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.ProfBonus != 0 {
		model.ProfBonus = request.ProfBonus
	}
	if request.SpellKnown != 0 {
		model.SpellKnown = request.SpellKnown
	}
	if request.SpellSlot1 != 0 {
		model.SpellSlot1 = request.SpellSlot1
	}
	if request.SpellSlot2 != 0 {
		model.SpellSlot2 = request.SpellSlot2
	}
	if request.SpellSlot3 != 0 {
		model.SpellSlot3 = request.SpellSlot3
	}
	if request.SpellSlot4 != 0 {
		model.SpellSlot4 = request.SpellSlot4
	}
	if request.SpellSlot5 != 0 {
		model.SpellSlot5 = request.SpellSlot5
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

func RangerTableHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.RangerTable

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
