package controllers

import (
	"strconv"

	inputs "github.com/fikrimohammad/ficree-api/app/inputs/users"
	"github.com/fikrimohammad/ficree-api/app/interfaces"
	"github.com/fikrimohammad/ficree-api/app/presenters"
	"github.com/fikrimohammad/ficree-api/app/services"
	"github.com/gofiber/fiber/v2"
)

// UsersController is a controller to handle APIs for User
type UsersController struct {
	svc interfaces.IUserService
}

// NewUsersController is a function to initialize UsersController
func NewUsersController(svc *services.UserService) UsersController {
	return UsersController{svc: svc}
}

// All is an API to fetch all active users
func (c *UsersController) All(ctx *fiber.Ctx) error {
	users, err := c.svc.All()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	results := []map[string]interface{}{}
	for _, user := range users {
		result := presenters.NewUserPresenter(user, "minimal_format").Result()
		results = append(results, result)
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": results,
	})
}

// Show is an API to find an user by ID
func (c *UsersController) Show(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	user, err := c.svc.Show(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewUserPresenter(user, "detail_format").Result(),
	})
}

// Create is an API to create an user
func (c *UsersController) Create(ctx *fiber.Ctx) error {
	input, inputErr := inputs.NewUserCreateInput(ctx)
	if inputErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	user, err := c.svc.Create(input)
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"data": presenters.NewUserPresenter(user, "detail_format").Result(),
	})
}

// Update is an API to update an user
func (c *UsersController) Update(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	input, inputErr := inputs.NewUserUpdateInput(ctx)
	if inputErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	user, err := c.svc.Update(id, input)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewUserPresenter(user, "detail_format").Result(),
	})
}

// Destroy is an API to destroy an user
func (c *UsersController) Destroy(ctx *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(ctx.Params("id"))
	if parseIntErr != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	user, err := c.svc.Destroy(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": presenters.NewUserPresenter(user, "detail_format").Result(),
	})
}
