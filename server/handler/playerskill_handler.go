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

func PlayerSkillHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.PlayerSkillProf
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func PlayerSkillHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.PlayerSkillCreateRequest)
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

	model := model.PlayerSkillProf{
		PlayerId: request.PlayerId,
		SkillId:  request.SkillId,
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

func PlayerSkillHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.PlayerSkillProf
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.PlayerSkillResponse{
		Id:       model.Id,
		PlayerId: model.PlayerId,
		SkillId:  model.SkillId,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func PlayerSkillHandlerGetByPlayerId(ctx *fiber.Ctx) error {
	playerId := ctx.Params("player_id")

	var models []model.PlayerSkillProf
	err := config.DB.Where("player_id = ?", playerId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.PlayerSkillResponse
	for _, model := range models {
		responses = append(responses, response.PlayerSkillResponse{
			Id:       model.Id,
			PlayerId: model.PlayerId,
			SkillId:  model.SkillId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func PlayerSkillHandlerGetBySkillId(ctx *fiber.Ctx) error {
	skillId := ctx.Params("skill_id")

	var models []model.PlayerSkillProf
	err := config.DB.Where("skill_id = ?", skillId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.PlayerSkillResponse
	for _, model := range models {
		responses = append(responses, response.PlayerSkillResponse{
			Id:       model.Id,
			PlayerId: model.PlayerId,
			SkillId:  model.SkillId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func PlayerSkillHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.PlayerSkillUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.PlayerSkillProf

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
	if request.SkillId != 0 {
		model.SkillId = request.SkillId
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

func PlayerSkillHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.PlayerSkillProf

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
