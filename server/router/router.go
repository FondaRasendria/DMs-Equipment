package router

import (
	"fmt"
	"golang-crud/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	raceController *controller.RaceController,
	subraceController *controller.SubraceController,
	traitController *controller.TraitController,
	subtraitController *controller.SubtraitController,
	monsterController *controller.MonsterController,
	monsterTraitController *controller.MonsterTraitController,
	actionController *controller.ActionController) *httprouter.Router {
	router := httprouter.New()

	corsMiddleware := func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9999")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			next(w, r, p)
		}
	}

	router.GET("/", corsMiddleware(homeHandler))

	router.GET("/api/race", corsMiddleware(raceController.SelectAll))
	router.GET("/api/race/id/:raceId", corsMiddleware(raceController.SelectById))
	router.POST("/api/race", corsMiddleware(raceController.Create))
	router.PUT("/api/race/:raceId", corsMiddleware(raceController.Update))
	router.DELETE("/api/race/:raceId", corsMiddleware(raceController.Delete))
	router.GET("/api/race/name/:name", corsMiddleware(raceController.SelectByName))

	router.GET("/api/subrace", corsMiddleware(subraceController.SelectAll))
	router.GET("/api/subrace/id/:subraceId", corsMiddleware(subraceController.SelectById))
	router.POST("/api/subrace", corsMiddleware(subraceController.Create))
	router.PUT("/api/subrace/:subraceId", corsMiddleware(subraceController.Update))
	router.DELETE("/api/subrace/:subraceId", corsMiddleware(subraceController.Delete))
	router.GET("/api/subrace/parent/:parentId", corsMiddleware(subraceController.SelectByParentId))

	router.GET("/api/trait", corsMiddleware(traitController.SelectAll))
	router.GET("/api/trait/id/:traitId", corsMiddleware(traitController.SelectById))
	router.POST("/api/trait", corsMiddleware(traitController.Create))
	router.PUT("/api/trait/:traitId", corsMiddleware(traitController.Update))
	router.DELETE("/api/trait/:traitId", corsMiddleware(traitController.Delete))
	router.GET("/api/trait/parent/:parentId", corsMiddleware(traitController.SelectByParentId))

	router.GET("/api/subtrait", corsMiddleware(subtraitController.SelectAll))
	router.GET("/api/subtrait/id/:subtraitId", corsMiddleware(subtraitController.SelectById))
	router.POST("/api/subtrait", corsMiddleware(subtraitController.Create))
	router.PUT("/api/subtrait/:subtraitId", corsMiddleware(subtraitController.Update))
	router.DELETE("/api/subtrait/:subtraitId", corsMiddleware(subtraitController.Delete))
	router.GET("/api/subtrait/parent/:parentId", corsMiddleware(subtraitController.SelectByParentId))

	router.GET("/api/monster", corsMiddleware(monsterController.SelectAll))
	router.GET("/api/monster/id/:monsterId", corsMiddleware(monsterController.SelectById))
	router.POST("/api/monster", corsMiddleware(monsterController.Create))
	router.PUT("/api/monster/:monsterId", corsMiddleware(monsterController.Update))
	router.DELETE("/api/monster/:monsterId", corsMiddleware(monsterController.Delete))
	router.GET("/api/monster/name/:name", corsMiddleware(monsterController.SelectByName))
	router.GET("/api/monster/source/:source", corsMiddleware(monsterController.SelectBySource))

	router.GET("/api/monstertrait", corsMiddleware(monsterTraitController.SelectAll))
	router.GET("/api/monstertrait/id/:mtId", corsMiddleware(monsterTraitController.SelectById))
	router.POST("/api/monstertrait", corsMiddleware(monsterTraitController.Create))
	router.PUT("/api/monstertrait/:mtId", corsMiddleware(monsterTraitController.Update))
	router.DELETE("/api/monstertrait/:mtId", corsMiddleware(monsterTraitController.Delete))
	router.GET("/api/monstertrait/parent/:parentId", corsMiddleware(monsterTraitController.SelectByParentId))

	router.GET("/api/action", corsMiddleware(actionController.SelectAll))
	router.GET("/api/action/id/:actionId", corsMiddleware(actionController.SelectById))
	router.POST("/api/action", corsMiddleware(actionController.Create))
	router.PUT("/api/action/:actionId", corsMiddleware(actionController.Update))
	router.DELETE("/api/action/:actionId", corsMiddleware(actionController.Delete))
	router.GET("/api/action/parent/:parentId", corsMiddleware(actionController.SelectByParentId))

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome Home")
}
