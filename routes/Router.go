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
	registerExperiencesAPI(api)
	registerEducationsAPI(api)
	registerUploadsAPI(api)
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

func registerExperiencesAPI(api fiber.Router) {
	controller := ServiceContainer().InjectExperiencesController()
	experiencesAPI := api.Group("/experiences")
	experiencesAPI.Get("/", controller.All)
	experiencesAPI.Get("/:id", controller.Show)
	experiencesAPI.Post("/", controller.Create)
	experiencesAPI.Patch("/:id", controller.Update)
	experiencesAPI.Delete("/:id", controller.Destroy)
}

func registerEducationsAPI(api fiber.Router) {
	controller := ServiceContainer().InjectEducationsController()
	educationsAPI := api.Group("/educations")
	educationsAPI.Get("/", controller.All)
	educationsAPI.Get("/:id", controller.Show)
	educationsAPI.Post("/", controller.Create)
	educationsAPI.Patch("/:id", controller.Update)
	educationsAPI.Delete("/:id", controller.Destroy)
}

func registerUploadsAPI(api fiber.Router) {
	controller := ServiceContainer().InjectUploadsController()
	uploadsAPI := api.Group("/uploads")
	uploadsAPI.Get("/presign", controller.Presign)
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
