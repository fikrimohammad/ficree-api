package controllers

import (
	"strconv"

	inputs "github.com/fikrimohammad/ficree-api/app/inputs/skills"
	"github.com/fikrimohammad/ficree-api/app/interfaces"
	"github.com/fikrimohammad/ficree-api/app/presenters"
	"github.com/fikrimohammad/ficree-api/app/services"
	"github.com/gofiber/fiber/v2"
)

// SkillsController is a controller to handle APIs for Skill
type SkillsController struct {
	svc interfaces.ISkillService
}

// NewSkillsController is a function to initialize SkillsController
func NewSkillsController(svc *services.SkillService) SkillsController {
	return SkillsController{svc: svc}
}

// All is an API to fetch all active skills
func (c *SkillsController) All(ctx *fiber.Ctx) error {
	skills, err := c.svc.All()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	results := []map[string]interface{}{}
	for _, skill := range skills {
		result := presenters.NewSkillPresenter(skill, "").Result()
		results = append(results, result)
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": results,
	})
}

// Show is an API to find an skill by ID
func (c *SkillsController) Show(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	skill, err := c.svc.Show(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewSkillPresenter(skill, "").Result(),
	})
}

// Create is an API to create an skill
func (c *SkillsController) Create(ctx *fiber.Ctx) error {
	input, inputErr := inputs.NewSkillCreateInput(ctx)
	if inputErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	skill, err := c.svc.Create(input)
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"data": presenters.NewSkillPresenter(skill, "").Result(),
	})
}

// Update is an API to update an skill
func (c *SkillsController) Update(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	input, inputErr := inputs.NewSkillUpdateInput(ctx)
	if inputErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	skill, err := c.svc.Update(id, input)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewSkillPresenter(skill, "").Result(),
	})
}

// Destroy is an API to destroy an skill
func (c *SkillsController) Destroy(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	skill, err := c.svc.Destroy(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewSkillPresenter(skill, "").Result(),
	})
}
