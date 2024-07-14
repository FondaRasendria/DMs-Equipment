package main

import (
	"fmt"
	"golang-crud/config"
	"golang-crud/migration"
	"golang-crud/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Printf("Start Server")

	config.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
		AllowOrigins: "*",
		AllowMethods: "*",
	}))
	router.RouteInit(app)

	app.Listen(":8888")
}
