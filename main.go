package main

import (
	"log"

	"github.com/fikrimohammad/ficree-api/common/apiresponse"
	"github.com/fikrimohammad/ficree-api/config"
	"github.com/fikrimohammad/ficree-api/infrastructure/database"
	"github.com/fikrimohammad/ficree-api/infrastructure/storage"
	"github.com/fikrimohammad/ficree-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func defaultErrorHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}
	return nil
}

func main() {
	app := fiber.New(
		fiber.Config{ErrorHandler: defaultErrorHandler},
	)

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	config.Load()
	database.Connect()
	storage.InitAWSInstance()

	router := routes.AppRouter()
	router.RegisterAPI(app)

	log.Fatal(app.Listen(":3000"))
}
