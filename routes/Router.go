package routes

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

// IAppRouter is an interface for AppRouter
type IAppRouter interface {
	RegisterAPI(app *fiber.App)
}

// RegisterAPI is a function to initialize API routes
func (router *router) RegisterAPI(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("")
	})
	registerUsersAPI(api)
	registerSkillsAPI(api)
}

func registerUsersAPI(api fiber.Router) {
	controller := ServiceContainer().InjectUsersController()
	usersAPI := api.Group("/users")
	usersAPI.Get("/", controller.All)
	usersAPI.Get("/:id", controller.Show)
	usersAPI.Post("/", controller.Create)
	usersAPI.Patch("/:id", controller.Update)
	usersAPI.Delete("/:id", controller.Destroy)
}

func registerSkillsAPI(api fiber.Router) {
	controller := ServiceContainer().InjectSkillsController()
	skillsAPI := api.Group("/skills")
	skillsAPI.Get("/", controller.All)
	skillsAPI.Get("/:id", controller.Show)
	skillsAPI.Post("/", controller.Create)
	skillsAPI.Patch("/:id", controller.Update)
	skillsAPI.Delete("/:id", controller.Destroy)
}

type router struct{}

var (
	r          *router
	routerOnce sync.Once
)

// AppRouter is a function to initialize IAppRouter instance
func AppRouter() IAppRouter {
	if r == nil {
		routerOnce.Do(func() {
			r = &router{}
		})
	}
	return r
}
