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

func PlayerHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Player
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func PlayerHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.PlayerCreateRequest)
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

	model := model.Player{
		PlayerName:        request.PlayerName,
		CharacterName:     request.CharacterName,
		RaceId:            request.RaceId,
		ClassId:           request.ClassId,
		SubclassId:        request.SubclassId,
		BackgroundId:      request.BackgroundId,
		XP:                request.XP,
		Level:             request.Level,
		AC:                request.AC,
		MaxHP:             request.MaxHP,
		HitDice:           request.HitDice,
		SpellSaveDC:       request.SpellSaveDC,
		SpellBonus:        request.SpellBonus,
		Speed:             request.Speed,
		STR:               request.STR,
		DEX:               request.DEX,
		COS:               request.COS,
		INT:               request.INT,
		WIS:               request.WIS,
		CHA:               request.CHA,
		PassivePerception: request.PassivePerception,
		PassiveInsight:    request.PassiveInsight,
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

func PlayerHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Player
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.PlayerResponse{
		Id:                model.Id,
		PlayerName:        model.PlayerName,
		CharacterName:     model.CharacterName,
		RaceId:            model.RaceId,
		ClassId:           model.ClassId,
		SubclassId:        model.SubclassId,
		BackgroundId:      model.BackgroundId,
		XP:                model.XP,
		Level:             model.Level,
		AC:                model.AC,
		MaxHP:             model.MaxHP,
		HitDice:           model.HitDice,
		SpellSaveDC:       model.SpellSaveDC,
		SpellBonus:        model.SpellBonus,
		Speed:             model.Speed,
		STR:               model.STR,
		DEX:               model.DEX,
		COS:               model.COS,
		INT:               model.INT,
		WIS:               model.WIS,
		CHA:               model.CHA,
		PassivePerception: model.PassivePerception,
		PassiveInsight:    model.PassiveInsight,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func PlayerHandlerGetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var models []model.Player
	err := config.DB.Where("character_name = ?", name).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.PlayerResponse
	for _, model := range models {
		responses = append(responses, response.PlayerResponse{
			Id:                model.Id,
			PlayerName:        model.PlayerName,
			CharacterName:     model.CharacterName,
			RaceId:            model.RaceId,
			ClassId:           model.ClassId,
			SubclassId:        model.SubclassId,
			BackgroundId:      model.BackgroundId,
			XP:                model.XP,
			Level:             model.Level,
			AC:                model.AC,
			MaxHP:             model.MaxHP,
			HitDice:           model.HitDice,
			SpellSaveDC:       model.SpellSaveDC,
			SpellBonus:        model.SpellBonus,
			Speed:             model.Speed,
			STR:               model.STR,
			DEX:               model.DEX,
			COS:               model.COS,
			INT:               model.INT,
			WIS:               model.WIS,
			CHA:               model.CHA,
			PassivePerception: model.PassivePerception,
			PassiveInsight:    model.PassiveInsight,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func PlayerHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.PlayerUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Player

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.PlayerName != "" {
		model.PlayerName = request.PlayerName
	}
	if request.CharacterName != "" {
		model.CharacterName = request.CharacterName
	}
	if request.RaceId != 0 {
		model.RaceId = request.RaceId
	}
	if request.ClassId != 0 {
		model.ClassId = request.ClassId
	}
	if request.BackgroundId != 0 {
		model.BackgroundId = request.BackgroundId
	}
	if request.XP != 0 {
		model.XP = request.XP
	}
	if request.Level != 0 {
		model.Level = request.Level
	}
	if request.AC != 0 {
		model.AC = request.AC
	}
	if request.MaxHP != 0 {
		model.MaxHP = request.MaxHP
	}
	if request.HitDice != "" {
		model.HitDice = request.HitDice
	}
	if request.SpellSaveDC != 0 {
		model.SpellSaveDC = request.SpellSaveDC
	}
	if request.SpellBonus != 0 {
		model.SpellBonus = request.SpellBonus
	}
	if request.Speed != 0 {
		model.Speed = request.Speed
	}
	if request.STR != 0 {
		model.STR = request.STR
	}
	if request.DEX != 0 {
		model.DEX = request.DEX
	}
	if request.COS != 0 {
		model.COS = request.COS
	}
	if request.INT != 0 {
		model.INT = request.INT
	}
	if request.WIS != 0 {
		model.WIS = request.WIS
	}
	if request.CHA != 0 {
		model.CHA = request.CHA
	}
	if request.PassivePerception != 0 {
		model.PassivePerception = request.PassivePerception
	}
	if request.PassiveInsight != 0 {
		model.PassiveInsight = request.PassiveInsight
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

func PlayerHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Player

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
