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

func MonkTableHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.MonkTable
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func MonkTableHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.MonkTableCreateRequest)
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

	model := model.MonkTable{
		ProfBonus:       request.ProfBonus,
		MartialArt:      request.MartialArt,
		KiPoint:         request.KiPoint,
		UnarmedMovement: request.UnarmedMovement,
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

func MonkTableHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.MonkTable
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.MonkTableResponse{
		Id:              model.Id,
		ProfBonus:       model.ProfBonus,
		MartialArt:      model.MartialArt,
		KiPoint:         model.KiPoint,
		UnarmedMovement: model.UnarmedMovement,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func MonkTableHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.MonkTableUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.MonkTable

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
	if request.MartialArt != "" {
		model.MartialArt = request.MartialArt
	}
	if request.KiPoint != 0 {
		model.KiPoint = request.KiPoint
	}
	if request.UnarmedMovement != "" {
		model.UnarmedMovement = request.UnarmedMovement
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

func MonkTableHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.MonkTable

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
