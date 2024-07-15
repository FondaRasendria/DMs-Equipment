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

func MonsterSkillProfHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.MonsterSkillProf
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func MonsterSkillProfHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.MonsterSkillProfCreateRequest)
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

	model := model.MonsterSkillProf{
		MonsterId: request.MonsterId,
		SkillId:   request.SkillId,
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

func MonsterSkillProfHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.MonsterSkillProf
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.MonsterSkillProfResponse{
		Id:        model.Id,
		MonsterId: model.MonsterId,
		SkillId:   model.SkillId,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func MonsterSkillProfHandlerGetByMonsterId(ctx *fiber.Ctx) error {
	monsterId := ctx.Params("monster_id")

	var models []model.MonsterSkillProf
	err := config.DB.Where("monster_id = ?", monsterId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.MonsterSkillProfResponse
	for _, model := range models {
		responses = append(responses, response.MonsterSkillProfResponse{
			Id:        model.Id,
			MonsterId: model.MonsterId,
			SkillId:   model.SkillId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func MonsterSkillProfHandlerGetBySkillId(ctx *fiber.Ctx) error {
	skillId := ctx.Params("skill_id")

	var models []model.MonsterSkillProf
	err := config.DB.Where("skill_id = ?", skillId).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.MonsterSkillProfResponse
	for _, model := range models {
		responses = append(responses, response.MonsterSkillProfResponse{
			Id:        model.Id,
			MonsterId: model.MonsterId,
			SkillId:   model.SkillId,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func MonsterSkillProfHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.MonsterSkillProfUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.MonsterSkillProf

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.MonsterId != 0 {
		model.MonsterId = request.MonsterId
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

func MonsterSkillProfHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.MonsterSkillProf

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
