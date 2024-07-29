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

func ArtificerTableHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.ArtificerTable
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func ArtificerTableHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.ArtificerTableCreateRequest)
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

	model := model.ArtificerTable{
		ProfBonus:     request.ProfBonus,
		InfusionKnown: request.InfusionKnown,
		InfusedItem:   request.InfusedItem,
		CantripKnown:  request.CantripKnown,
		SpellSlot1:    request.SpellSlot1,
		SpellSlot2:    request.SpellSlot2,
		SpellSlot3:    request.SpellSlot3,
		SpellSlot4:    request.SpellSlot4,
		SpellSlot5:    request.SpellSlot5,
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

func ArtificerTableHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.ArtificerTable
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.ArtificerTableResponse{
		Id:            model.Id,
		ProfBonus:     model.ProfBonus,
		InfusionKnown: model.InfusionKnown,
		InfusedItem:   model.InfusedItem,
		CantripKnown:  model.CantripKnown,
		SpellSlot1:    model.SpellSlot1,
		SpellSlot2:    model.SpellSlot2,
		SpellSlot3:    model.SpellSlot3,
		SpellSlot4:    model.SpellSlot4,
		SpellSlot5:    model.SpellSlot5,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func ArtificerTableHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.ArtificerTableUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.ArtificerTable

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
	if request.InfusionKnown != 0 {
		model.InfusionKnown = request.InfusionKnown
	}
	if request.InfusedItem != 0 {
		model.InfusedItem = request.InfusedItem
	}
	if request.CantripKnown != 0 {
		model.CantripKnown = request.CantripKnown
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

func ArtificerTableHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.ArtificerTable

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
