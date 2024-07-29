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

func BloodHunterTableHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.BloodHunterTable
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func BloodHunterTableHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.BloodHunterTableCreateRequest)
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

	model := model.BloodHunterTable{
		ProfBonus:       request.ProfBonus,
		HemocraftDie:    request.HemocraftDie,
		BloodCurseKnown: request.BloodCurseKnown,
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

func BloodHunterTableHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.BloodHunterTable
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.BloodHunterTableResponse{
		Id:              model.Id,
		ProfBonus:       model.ProfBonus,
		HemocraftDie:    model.HemocraftDie,
		BloodCurseKnown: model.BloodCurseKnown,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func BloodHunterTableHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.BloodHunterTableUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.BloodHunterTable

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
	if request.HemocraftDie != "" {
		model.HemocraftDie = request.HemocraftDie
	}
	if request.BloodCurseKnown != 0 {
		model.BloodCurseKnown = request.BloodCurseKnown
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

func BloodHunterTableHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.BloodHunterTable

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
