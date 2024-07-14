package migration

import (
	"golang-crud/config"
	"golang-crud/model"
)

func RunMigration() {
	config.DB.AutoMigrate(&model.Action{})
	config.DB.AutoMigrate(&model.Monster{})
	config.DB.AutoMigrate(&model.MonsterTrait{})
	config.DB.AutoMigrate(&model.Race{})
	config.DB.AutoMigrate(&model.Subrace{})
	config.DB.AutoMigrate(&model.Subtrait{})
	config.DB.AutoMigrate(&model.Trait{})

	config.DB.AutoMigrate(&model.ConditionType{})
	config.DB.AutoMigrate(&model.DamageType{})
	config.DB.AutoMigrate(&model.Language{})
	config.DB.AutoMigrate(&model.SavingThrow{})
	config.DB.AutoMigrate(&model.SkillType{})
}
