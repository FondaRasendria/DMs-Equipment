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

func ClassHandlerGetAll(ctx *fiber.Ctx) error {
	var models []model.Class
	result := config.DB.Find(&models)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(models)
}

func ClassHandlerCreate(ctx *fiber.Ctx) error {
	request := new(request.ClassCreateRequest)
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

	model := model.Class{
		Name:           request.Name,
		Description:    request.Description,
		Multiclass:     request.Multiclass,
		TableId:        request.TableId,
		HitDice:        request.HitDice,
		HitPointFirst:  request.HitPointFirst,
		HitPointHigher: request.HitPointHigher,
		ArmorProf:      request.ArmorProf,
		WeaponProf:     request.WeaponProf,
		ToolProf:       request.ToolProf,
		SavingProf:     request.SavingProf,
		SkillProf:      request.SkillProf,
		Equipment:      request.Equipment,
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

func ClassHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var model model.Class
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	response := response.ClassResponse{
		Id:             model.Id,
		Name:           model.Name,
		Description:    model.Description,
		Multiclass:     model.Multiclass,
		TableId:        model.TableId,
		HitDice:        model.HitDice,
		HitPointFirst:  model.HitPointFirst,
		HitPointHigher: model.HitPointHigher,
		ArmorProf:      model.ArmorProf,
		WeaponProf:     model.WeaponProf,
		ToolProf:       model.ToolProf,
		SavingProf:     model.SavingProf,
		SkillProf:      model.SkillProf,
		Equipment:      model.Equipment,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func ClassHandlerGetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var models []model.Class
	err := config.DB.Where("name = ?", name).Find(&models).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	var responses []response.ClassResponse
	for _, model := range models {
		responses = append(responses, response.ClassResponse{
			Id:             model.Id,
			Name:           model.Name,
			Description:    model.Description,
			Multiclass:     model.Multiclass,
			TableId:        model.TableId,
			HitDice:        model.HitDice,
			HitPointFirst:  model.HitPointFirst,
			HitPointHigher: model.HitPointHigher,
			ArmorProf:      model.ArmorProf,
			WeaponProf:     model.WeaponProf,
			ToolProf:       model.ToolProf,
			SavingProf:     model.SavingProf,
			SkillProf:      model.SkillProf,
			Equipment:      model.Equipment,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    responses,
	})
}

func ClassHandlerUpdate(ctx *fiber.Ctx) error {
	request := new(request.ClassUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var model model.Class

	id := ctx.Params("id")
	err := config.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	if request.Name != "" {
		model.Name = request.Name
	}
	if request.Description != "" {
		model.Description = request.Description
	}
	if request.Multiclass != "" {
		model.Multiclass = request.Multiclass
	}
	if request.TableId != 0 {
		model.TableId = request.TableId
	}
	if request.HitDice != "" {
		model.HitDice = request.HitDice
	}
	if request.HitPointFirst != "" {
		model.HitPointFirst = request.HitPointFirst
	}
	if request.HitPointHigher != "" {
		model.HitPointHigher = request.HitPointHigher
	}
	if request.ArmorProf != "" {
		model.ArmorProf = request.ArmorProf
	}
	if request.WeaponProf != "" {
		model.WeaponProf = request.WeaponProf
	}
	if request.ToolProf != "" {
		model.ToolProf = request.ToolProf
	}
	if request.SavingProf != "" {
		model.SavingProf = request.SavingProf
	}
	if request.SkillProf != "" {
		model.SkillProf = request.SkillProf
	}
	if request.Equipment != "" {
		model.Equipment = request.Equipment
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

func ClassHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var model model.Class

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
