package controllers

import (
	"strconv"

	"github.com/fikrimohammad/ficree-api/app/inputs"
	"github.com/fikrimohammad/ficree-api/app/interfaces"
	"github.com/fikrimohammad/ficree-api/app/presenters"
	"github.com/fikrimohammad/ficree-api/app/services"
	"github.com/gofiber/fiber/v2"
)

// EducationsController is a controller to handle APIs for Education
type EducationsController struct {
	svc interfaces.IEducationService
}

// NewEducationsController is a function to initialize EducationsController
func NewEducationsController(svc *services.EducationService) EducationsController {
	return EducationsController{svc: svc}
}

// All is an API to fetch all active educations
func (c *EducationsController) All(ctx *fiber.Ctx) error {
	educations, err := c.svc.All()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	results := []map[string]interface{}{}
	for _, education := range educations {
		result := presenters.NewEducationPresenter(education, "").Result()
		results = append(results, result)
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": results,
	})
}

// Show is an API to find an education by ID
func (c *EducationsController) Show(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	education, err := c.svc.Show(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewEducationPresenter(education, "").Result(),
	})
}

// Create is an API to create an education
func (c *EducationsController) Create(ctx *fiber.Ctx) error {
	input, inputErr := inputs.NewEducationCreateInput(ctx)
	if inputErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	education, err := c.svc.Create(input)
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"data": presenters.NewEducationPresenter(education, "").Result(),
	})
}

// Update is an API to update an education
func (c *EducationsController) Update(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	input, inputErr := inputs.NewEducationUpdateInput(ctx)
	if inputErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	education, err := c.svc.Update(id, input)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewEducationPresenter(education, "").Result(),
	})
}

// Destroy is an API to destroy an education
func (c *EducationsController) Destroy(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	education, err := c.svc.Destroy(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewEducationPresenter(education, "").Result(),
	})
}
