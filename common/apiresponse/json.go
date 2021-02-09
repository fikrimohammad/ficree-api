package apiresponse

import "github.com/gofiber/fiber/v2"

// RenderJSONSuccess is a helper function to render JSON for successful operations
func RenderJSONSuccess(ctx *fiber.Ctx, statusCode int, data interface{}) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"data": data,
	})
}

// RenderJSONError is a helper function to render error into JSON
func RenderJSONError(ctx *fiber.Ctx, statusCode int, err error) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"error": err.Error(),
	})
}
