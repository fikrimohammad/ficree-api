package controllers

import (
	"strconv"

	"github.com/fikrimohammad/ficree-api/app/inputs"
	"github.com/fikrimohammad/ficree-api/app/interfaces"
	"github.com/fikrimohammad/ficree-api/app/presenters"
	"github.com/fikrimohammad/ficree-api/app/services"
	"github.com/gofiber/fiber/v2"
)

// ExperiencesController is a controller to handle APIs for Experience
type ExperiencesController struct {
	svc interfaces.IExperienceService
}

// NewExperiencesController is a function to initialize ExperiencesController
func NewExperiencesController(svc *services.ExperienceService) ExperiencesController {
	return ExperiencesController{svc: svc}
}

// All is an API to fetch all active experiences
func (c *ExperiencesController) All(ctx *fiber.Ctx) error {
	experiences, err := c.svc.All()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	results := []map[string]interface{}{}
	for _, experience := range experiences {
		result := presenters.NewExperiencePresenter(experience, "").Result()
		results = append(results, result)
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": results,
	})
}

// Show is an API to find an experience by ID
func (c *ExperiencesController) Show(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	experience, err := c.svc.Show(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewExperiencePresenter(experience, "").Result(),
	})
}

// Create is an API to create an experience
func (c *ExperiencesController) Create(ctx *fiber.Ctx) error {
	input, inputErr := inputs.NewExperienceCreateInput(ctx)
	if inputErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	experience, err := c.svc.Create(input)
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"data": presenters.NewExperiencePresenter(experience, "").Result(),
	})
}

// Update is an API to update an experience
func (c *ExperiencesController) Update(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	input, inputErr := inputs.NewExperienceUpdateInput(ctx)
	if inputErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	experience, err := c.svc.Update(id, input)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewExperiencePresenter(experience, "").Result(),
	})
}

// Destroy is an API to destroy an experience
func (c *ExperiencesController) Destroy(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	experience, err := c.svc.Destroy(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewExperiencePresenter(experience, "").Result(),
	})
}
