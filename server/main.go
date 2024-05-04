package main

import (
	"fmt"
	"golang-crud/config"
	"golang-crud/controller"
	"golang-crud/helper"
	"golang-crud/repository"
	"golang-crud/router"
	"golang-crud/service"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Printf("Start Server")
	//database
	db := config.DatabaseConnection()

	//repository
	raceRepository := repository.NewRaceRepository(db)
	subraceRepository := repository.NewSubraceRepository(db)
	traitRepository := repository.NewTraitRepository(db)
	subtraitRepository := repository.NewSubtraitRepository(db)
	monsterRepository := repository.NewMonsterRepository(db)
	monsterTraitRepository := repository.NewMonsterTraitRepositoryImpl(db)
	actionRepository := repository.NewActionRepository(db)

	//service
	raceService := service.NewRaceServiceImpl(raceRepository)
	subraceService := service.NewSubraceServiceImpl(subraceRepository)
	traitService := service.NewTraitServiceImpl(traitRepository)
	subtraitService := service.NewSubtraitServiceImpl(subtraitRepository)
	monsterService := service.NewMonsterServiceImpl(monsterRepository)
	monsterTraitService := service.NewMonsterTraitServiceImpl(monsterTraitRepository)
	actionService := service.NewActionServiceImpl(actionRepository)

	//controller
	raceController := controller.NewRaceController(raceService)
	subraceController := controller.NewSubraceController(subraceService)
	traitController := controller.NewTraitController(traitService)
	subtraitController := controller.NewSubtraitController(subtraitService)
	monsterController := controller.NewMonsterController(monsterService)
	monsterTraitController := controller.NewMonsterTraitController(monsterTraitService)
	actionController := controller.NewActionController(actionService)

	//router
	routes := router.NewRouter(
		raceController,
		subraceController,
		traitController,
		subtraitController,
		monsterController,
		monsterTraitController,
		actionController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
