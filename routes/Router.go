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
	registerFilesAPI(api)
}

func registerUsersAPI(api fiber.Router) {
	handler := ServiceContainer().InjectUserHandler()
	usersAPI := api.Group("/users")
	usersAPI.Get("/", handler.HandleListUsers)
	usersAPI.Get("/:id", handler.HandleShowUser)
	usersAPI.Post("/", handler.HandleCreateUser)
	usersAPI.Patch("/:id", handler.HandleUpdateUser)
	usersAPI.Delete("/:id", handler.HandleDestroyUser)
}

func registerFilesAPI(api fiber.Router) {
	handler := ServiceContainer().InjectFileHandler()
	filesAPI := api.Group("/files")
	filesAPI.Get("/presign", handler.HandleCreatePresignedURL)
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
