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

func MonsterHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Monster
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func MonsterHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.MonsterCreateRequest)
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

	model := model.Monster{
		Name:              request.Name,
		Type:              request.Type,
		Alignment:         request.Alignment,
		AC:                request.AC,
		FixedHP:           request.FixedHP,
		DiceHP:            request.DiceHP,
		Speed:             request.Speed,
		STR:               request.STR,
		DEX:               request.DEX,
		CON:               request.CON,
		INT:               request.INT,
		WIS:               request.WIS,
		CHA:               request.CHA,
		Senses:            request.Senses,
		PassivePerception: request.PassivePerception,
		CR:                request.CR,
		XPReward:          request.XPReward,
		Description:       request.Description,
		Source:            request.Source,
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

func MonsterHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Monster
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.MonsterResponse{
		Id:                model.Id,
		Name:              model.Name,
		Type:              model.Type,
		Alignment:         model.Alignment,
		AC:                model.AC,
		FixedHP:           model.FixedHP,
		DiceHP:            model.DiceHP,
		Speed:             model.Speed,
		STR:               model.STR,
		DEX:               model.DEX,
		CON:               model.CON,
		INT:               model.INT,
		WIS:               model.WIS,
		CHA:               model.CHA,
		Senses:            model.Senses,
		PassivePerception: model.PassivePerception,
		CR:                model.CR,
		XPReward:          model.XPReward,
		Description:       model.Description,
		Source:            model.Source,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func MonsterHandlerGetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var model model.Monster
	err := config.DB.Where("name = ?", name).Find(&model).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.MonsterResponse{
		Id:                model.Id,
		Name:              model.Name,
		Type:              model.Type,
		Alignment:         model.Alignment,
		AC:                model.AC,
		FixedHP:           model.FixedHP,
		DiceHP:            model.DiceHP,
		STR:               model.STR,
		DEX:               model.DEX,
		CON:               model.CON,
		INT:               model.INT,
		WIS:               model.WIS,
		CHA:               model.CHA,
		Senses:            model.Senses,
		PassivePerception: model.PassivePerception,
		CR:                model.CR,
		XPReward:          model.XPReward,
		Description:       model.Description,
		Source:            model.Source,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func MonsterHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.MonsterUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Monster

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	model.Name = request.Name
	model.Type = request.Type
	model.Alignment = request.Alignment
	model.AC = request.AC
	model.FixedHP = request.FixedHP
	model.DiceHP = request.DiceHP
	model.Speed = request.Speed
	model.STR = request.STR
	model.DEX = request.DEX
	model.CON = request.CON
	model.INT = request.INT
	model.WIS = request.WIS
	model.CHA = request.CHA
	model.Senses = request.Senses
	model.PassivePerception = request.PassivePerception
	model.CR = request.CR
	model.XPReward = request.XPReward
	model.Description = request.Description
	model.Source = request.Source

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

func MonsterHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Monster

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
