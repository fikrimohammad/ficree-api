package routes

import (
	"sync"

	"github.com/fikrimohammad/ficree-api/infrastructure/database"
	_file_handler "github.com/fikrimohammad/ficree-api/pkg/file/handler"
	_file_repo "github.com/fikrimohammad/ficree-api/pkg/file/repository"
	_file_svc "github.com/fikrimohammad/ficree-api/pkg/file/service"
	_user_handler "github.com/fikrimohammad/ficree-api/pkg/user/handler"
	_user_repo "github.com/fikrimohammad/ficree-api/pkg/user/repository"
	_user_svc "github.com/fikrimohammad/ficree-api/pkg/user/service"
)

// IServiceContainer represents ServiceContainer
type IServiceContainer interface {
	InjectUserHandler() _user_handler.UserHTTPHandler
	InjectFileHandler() _file_handler.FileHTTPHandler
}

type kernel struct{}

func (k *kernel) InjectUserHandler() _user_handler.UserHTTPHandler {
	db := database.Load()
	repo := _user_repo.NewSQLUserRepository(db)
	svc := _user_svc.NewUserService(repo)
	handler := _user_handler.NewUserHTTPHandler(svc)
	return handler
}

func (k *kernel) InjectFileHandler() _file_handler.FileHTTPHandler {
	repo := _file_repo.NewAWSFileRepository()
	svc := _file_svc.NewFileService(repo)
	handler := _file_handler.NewFileHTTPHandler(svc)
	return handler
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
