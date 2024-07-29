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

func WizardTableHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.WizardTable
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func WizardTableHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.WizardTableCreateRequest)
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

	model := model.WizardTable{
		ProfBonus:    request.ProfBonus,
		CantripKnown: request.CantripKnown,
		SpellSlot1:   request.SpellSlot1,
		SpellSlot2:   request.SpellSlot2,
		SpellSlot3:   request.SpellSlot3,
		SpellSlot4:   request.SpellSlot4,
		SpellSlot5:   request.SpellSlot5,
		SpellSlot6:   request.SpellSlot6,
		SpellSlot7:   request.SpellSlot7,
		SpellSlot8:   request.SpellSlot8,
		SpellSlot9:   request.SpellSlot9,
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

func WizardTableHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.WizardTable
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.WizardTableResponse{
		Id:           model.Id,
		ProfBonus:    model.ProfBonus,
		CantripKnown: model.CantripKnown,
		SpellSlot1:   model.SpellSlot1,
		SpellSlot2:   model.SpellSlot2,
		SpellSlot3:   model.SpellSlot3,
		SpellSlot4:   model.SpellSlot4,
		SpellSlot5:   model.SpellSlot5,
		SpellSlot6:   model.SpellSlot6,
		SpellSlot7:   model.SpellSlot7,
		SpellSlot8:   model.SpellSlot8,
		SpellSlot9:   model.SpellSlot9,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func WizardTableHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.WizardTableUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.WizardTable

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
	if request.SpellSlot6 != 0 {
		model.SpellSlot6 = request.SpellSlot6
	}
	if request.SpellSlot7 != 0 {
		model.SpellSlot7 = request.SpellSlot7
	}
	if request.SpellSlot8 != 0 {
		model.SpellSlot8 = request.SpellSlot8
	}
	if request.SpellSlot9 != 0 {
		model.SpellSlot9 = request.SpellSlot9
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

func WizardTableHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.WizardTable

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
