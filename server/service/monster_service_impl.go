package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
)

type MonsterServiceImpl struct {
	MonsterRepository repository.MonsterRepository
}

// Create implements MonsterService.
func (s *MonsterServiceImpl) Create(ctx context.Context, request request.MonsterCreateRequest) response.MonsterResponse {
	monster := model.Monster{
		Name:                  request.Name,
		Type:                  request.Type,
		Alignment:             request.Alignment,
		AC:                    request.AC,
		FixedHP:               request.FixedHP,
		DiceHP:                request.DiceHP,
		Speed:                 request.Speed,
		STR:                   request.STR,
		DEX:                   request.DEX,
		CON:                   request.CON,
		INT:                   request.INT,
		WIS:                   request.WIS,
		CHA:                   request.CHA,
		SavingThrows:          request.SavingThrows,
		Skills:                request.Skills,
		DamageImmunities:      request.DamageImmunities,
		DamageVulnerabilities: request.DamageVulnerabilities,
		DamageResistences:     request.DamageResistences,
		Senses:                request.Senses,
		Language:              request.Language,
		CR:                    request.CR,
		Description:           request.Description,
		Source:                request.Source,
	}
	s.MonsterRepository.Create(ctx, monster)
	return response.MonsterResponse(monster)
}

// Delete implements MonsterService.
func (s *MonsterServiceImpl) Delete(ctx context.Context, monsterId int) response.MonsterResponse {
	monster, err := s.MonsterRepository.SelectById(ctx, monsterId)
	helper.PanicIfError(err)
	s.MonsterRepository.Delete(ctx, monster.Id)
	return response.MonsterResponse(monster)
}

// SelectAll implements MonsterService.
func (s *MonsterServiceImpl) SelectAll(ctx context.Context) []response.MonsterResponse {
	monsters := s.MonsterRepository.SelectAll(ctx)

	var monstersResp []response.MonsterResponse

	for _, value := range monsters {
		monster := response.MonsterResponse{
			Id:                    value.Id,
			Name:                  value.Name,
			Type:                  value.Type,
			Alignment:             value.Alignment,
			AC:                    value.AC,
			FixedHP:               value.FixedHP,
			DiceHP:                value.DiceHP,
			Speed:                 value.Speed,
			STR:                   value.STR,
			DEX:                   value.DEX,
			CON:                   value.CON,
			INT:                   value.INT,
			WIS:                   value.WIS,
			CHA:                   value.CHA,
			SavingThrows:          value.SavingThrows,
			Skills:                value.Skills,
			DamageImmunities:      value.DamageImmunities,
			DamageVulnerabilities: value.DamageVulnerabilities,
			DamageResistences:     value.DamageResistences,
			Senses:                value.Senses,
			Language:              value.Language,
			CR:                    value.CR,
			Description:           value.Description,
			Source:                value.Source}
		monstersResp = append(monstersResp, monster)
	}

	return monstersResp
}

// SelectById implements MonsterService.
func (s *MonsterServiceImpl) SelectById(ctx context.Context, monsterId int) response.MonsterResponse {
	monster, err := s.MonsterRepository.SelectById(ctx, monsterId)
	helper.PanicIfError(err)
	return response.MonsterResponse(monster)
}

// SelectByName implements MonsterService.
func (s *MonsterServiceImpl) SelectByName(ctx context.Context, name string) response.MonsterResponse {
	monster, err := s.MonsterRepository.SelectByName(ctx, name)
	helper.PanicIfError(err)
	return response.MonsterResponse(monster)
}

// SelectBySource implements MonsterService.
func (s *MonsterServiceImpl) SelectBySource(ctx context.Context, source string) []response.MonsterResponse {
	monsters := s.MonsterRepository.SelectBySource(ctx, source)

	var monstersResp []response.MonsterResponse

	for _, value := range monsters {
		monster := response.MonsterResponse{
			Id:                    value.Id,
			Name:                  value.Name,
			Type:                  value.Type,
			Alignment:             value.Alignment,
			AC:                    value.AC,
			FixedHP:               value.FixedHP,
			DiceHP:                value.DiceHP,
			Speed:                 value.Speed,
			STR:                   value.STR,
			DEX:                   value.DEX,
			CON:                   value.CON,
			INT:                   value.INT,
			WIS:                   value.WIS,
			CHA:                   value.CHA,
			SavingThrows:          value.SavingThrows,
			Skills:                value.Skills,
			DamageImmunities:      value.DamageImmunities,
			DamageVulnerabilities: value.DamageVulnerabilities,
			DamageResistences:     value.DamageResistences,
			Senses:                value.Senses,
			Language:              value.Language,
			CR:                    value.CR,
			Description:           value.Description,
			Source:                value.Source}
		monstersResp = append(monstersResp, monster)
	}

	return monstersResp
}

// Update implements MonsterService.
func (s *MonsterServiceImpl) Update(ctx context.Context, request request.MonsterUpdateRequest) response.MonsterResponse {
	monster, err := s.MonsterRepository.SelectById(ctx, request.Id)
	helper.PanicIfError(err)

	monster.Name = request.Name
	monster.Type = request.Type
	monster.Alignment = request.Alignment
	monster.AC = request.AC
	monster.FixedHP = request.FixedHP
	monster.DiceHP = request.DiceHP
	monster.Speed = request.Speed
	monster.STR = request.STR
	monster.DEX = request.DEX
	monster.CON = request.CON
	monster.INT = request.INT
	monster.WIS = request.WIS
	monster.CHA = request.CHA
	monster.SavingThrows = request.SavingThrows
	monster.Skills = request.Skills
	monster.DamageImmunities = request.DamageImmunities
	monster.DamageVulnerabilities = request.DamageVulnerabilities
	monster.DamageResistences = request.DamageResistences
	monster.Senses = request.Senses
	monster.Language = request.Language
	monster.CR = request.CR
	monster.Description = request.Description
	monster.Source = request.Source
	s.MonsterRepository.Update(ctx, monster)
	return response.MonsterResponse(monster)
}

func NewMonsterServiceImpl(monsterRepository repository.MonsterRepository) MonsterService {
	return &MonsterServiceImpl{MonsterRepository: monsterRepository}
}
