package main

import (
	"log"

	"github.com/fikrimohammad/ficree-api/config"
	"github.com/fikrimohammad/ficree-api/database"
	"github.com/fikrimohammad/ficree-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	config.LoadAppConfig()

	database.Connect()
	database.Migrate()

	router := routes.AppRouter()
	router.RegisterAPI(app)

	log.Fatal(app.Listen(":3000"))
}
