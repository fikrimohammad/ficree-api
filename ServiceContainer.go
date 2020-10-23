package main

import (
	"sync"

	"github.com/fikrimohammad/ficree-api/app/controllers"
	"github.com/fikrimohammad/ficree-api/app/repositories"
	"github.com/fikrimohammad/ficree-api/app/services"
	"github.com/fikrimohammad/ficree-api/database"
)

// IServiceContainer represents ServiceContainer
type IServiceContainer interface {
	InjectUsersController() controllers.UsersController
}

type kernel struct{}

func (k *kernel) InjectUsersController() controllers.UsersController {
	dbConn := database.Instance()
	repo := repositories.NewUserRepository(dbConn)
	svc := services.NewUserService(repo)
	controller := controllers.NewUsersController(svc)
	return controller
}

var (
	k             *kernel
	containerOnce sync.Once
)

// ServiceContainer is a function to inject dependency
func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
