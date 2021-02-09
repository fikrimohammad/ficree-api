package main

import (
	"log"

	"github.com/fikrimohammad/ficree-api/config"
	"github.com/fikrimohammad/ficree-api/infrastructures/database"
	"github.com/fikrimohammad/ficree-api/infrastructures/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func defaultErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if err != nil {
		return ctx.Status(code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return nil
}

func main() {
	appConfig := fiber.Config{
		ErrorHandler: defaultErrorHandler,
	}
	app := fiber.New(appConfig)

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	config.Load()

	database.Connect()

	storage.InitAWSInstance()

	// router := routes.AppRouter()
	// router.RegisterAPI(app)

	log.Fatal(app.Listen(":3000"))
}
