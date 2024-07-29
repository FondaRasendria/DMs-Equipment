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

func BarbarianTableHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.BarbarianTable
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func BarbarianTableHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.BarbarianTableCreateRequest)
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

	model := model.BarbarianTable{
		ProfBonus:  request.ProfBonus,
		Rage:       request.Rage,
		RageDamage: request.RageDamage,
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

func BarbarianTableHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.BarbarianTable
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.BarbarianTableResponse{
		Id:         model.Id,
		ProfBonus:  model.ProfBonus,
		Rage:       model.Rage,
		RageDamage: model.RageDamage,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func BarbarianTableHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.BarbarianTableUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.BarbarianTable

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
	if request.Rage != 0 {
		model.Rage = request.Rage
	}
	if request.RageDamage != 0 {
		model.RageDamage = request.RageDamage
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

func BarbarianTableHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.BarbarianTable

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
