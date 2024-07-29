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

func WarlockTableHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.WarlockTable
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func WarlockTableHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.WarlockTableCreateRequest)
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

	model := model.WarlockTable{
		ProfBonus:       request.ProfBonus,
		CantripKnown:    request.CantripKnown,
		SpellKnown:      request.SpellKnown,
		SpellSlot:       request.SpellSlot,
		SlotLevel:       request.SlotLevel,
		InvocationKnown: request.InvocationKnown,
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

func WarlockTableHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.WarlockTable
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.WarlockTableResponse{
		Id:              model.Id,
		ProfBonus:       model.ProfBonus,
		CantripKnown:    model.CantripKnown,
		SpellKnown:      model.SpellKnown,
		SpellSlot:       model.SpellSlot,
		SlotLevel:       model.SlotLevel,
		InvocationKnown: model.InvocationKnown,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func WarlockTableHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.WarlockTableUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.WarlockTable

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
	if request.CantripKnown != 0 {
		model.CantripKnown = request.CantripKnown
	}
	if request.SpellKnown != 0 {
		model.SpellKnown = request.SpellKnown
	}
	if request.SpellSlot != 0 {
		model.SpellSlot = request.SpellSlot
	}
	if request.SlotLevel != "" {
		model.SlotLevel = request.SlotLevel
	}
	if request.InvocationKnown != 0 {
		model.InvocationKnown = request.InvocationKnown
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

func WarlockTableHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.WarlockTable

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
