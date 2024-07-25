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

func PlayerConditionImmunHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.PlayerConditionImmunities
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func PlayerConditionImmunHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.PlayerConditionImmunCreateRequest)
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

	model := model.PlayerConditionImmunities{
		PlayerId:    request.PlayerId,
		ConditionId: request.ConditionId,
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

func PlayerConditionImmunHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.PlayerConditionImmunities
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.PlayerConditionImmunResponse{
		Id:          model.Id,
		PlayerId:    model.PlayerId,
		ConditionId: model.ConditionId,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func PlayerConditionImmunHandlerGetByPlayerId(ctx *fiber.Ctx) error {
	playerId := ctx.Params("player_id")

	var models []model.PlayerConditionImmunities
	err := config.DB.Where("player_id = ?", playerId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.PlayerConditionImmunResponse
	for _, model := range models {
		responses = append(responses, response.PlayerConditionImmunResponse{
			Id:          model.Id,
			PlayerId:    model.PlayerId,
			ConditionId: model.ConditionId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func PlayerConditionImmunHandlerGetByConditionId(ctx *fiber.Ctx) error {
	conditionId := ctx.Params("condition_id")

	var models []model.PlayerConditionImmunities
	err := config.DB.Where("condition_id = ?", conditionId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.PlayerConditionImmunResponse
	for _, model := range models {
		responses = append(responses, response.PlayerConditionImmunResponse{
			Id:          model.Id,
			PlayerId:    model.PlayerId,
			ConditionId: model.ConditionId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func PlayerConditionImmunHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.PlayerConditionImmunUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.PlayerConditionImmunities

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.PlayerId != 0 {
		model.PlayerId = request.PlayerId
	}
	if request.ConditionId != 0 {
		model.ConditionId = request.ConditionId
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

func PlayerConditionImmunHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.PlayerConditionImmunities

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
