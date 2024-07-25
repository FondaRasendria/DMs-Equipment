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

func PlayerDamageVulneHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.PlayerDamageVulnerabilities
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func PlayerDamageVulneHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.PlayerDamageVulneCreateRequest)
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

	model := model.PlayerDamageVulnerabilities{
		PlayerId: request.PlayerId,
		DamageId: request.DamageId,
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

func PlayerDamageVulneHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.PlayerDamageVulnerabilities
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.PlayerDamageVulneResponse{
		Id:       model.Id,
		PlayerId: model.PlayerId,
		DamageId: model.DamageId,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func PlayerDamageVulneHandlerGetByPlayerId(ctx *fiber.Ctx) error {
	playerId := ctx.Params("player_id")

	var models []model.PlayerDamageVulnerabilities
	err := config.DB.Where("player_id = ?", playerId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.PlayerDamageVulneResponse
	for _, model := range models {
		responses = append(responses, response.PlayerDamageVulneResponse{
			Id:       model.Id,
			PlayerId: model.PlayerId,
			DamageId: model.DamageId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func PlayerDamageVulneHandlerGetByDamageId(ctx *fiber.Ctx) error {
	damageId := ctx.Params("damage_id")

	var models []model.PlayerDamageVulnerabilities
	err := config.DB.Where("damage_id = ?", damageId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.PlayerDamageVulneResponse
	for _, model := range models {
		responses = append(responses, response.PlayerDamageVulneResponse{
			Id:       model.Id,
			PlayerId: model.PlayerId,
			DamageId: model.DamageId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func PlayerDamageVulneHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.PlayerDamageVulneUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.PlayerDamageVulnerabilities

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
	if request.DamageId != 0 {
		model.DamageId = request.DamageId
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

func PlayerDamageVulneHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.PlayerDamageVulnerabilities

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
