package controllers

import (
	"github.com/fikrimohammad/ficree-api/app/inputs"
	"github.com/fikrimohammad/ficree-api/app/interfaces"
	"github.com/fikrimohammad/ficree-api/app/presenters"
	"github.com/fikrimohammad/ficree-api/app/services"
	"github.com/gofiber/fiber/v2"
)

type UploadsController struct {
	svc interfaces.IUploadService
}

func NewUploadsController(svc *services.UploadService) UploadsController {
	return UploadsController{svc: svc}
}

func (c *UploadsController) Presign(ctx *fiber.Ctx) error {
	input, inputErr := inputs.NewUploadPresignInput(ctx)
	if inputErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	file, err := c.svc.BuildPresignedURL(input.Output())
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewUploadPresenter(file, "").Result(),
	})
}
