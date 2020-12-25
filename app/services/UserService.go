package services

import (
	"github.com/fikrimohammad/ficree-api/app/inputs"
	"github.com/fikrimohammad/ficree-api/app/interfaces"
	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/fikrimohammad/ficree-api/app/repositories"
)

// UserService is a struct to wrap service for User
type UserService struct {
	repo interfaces.IUserRepository
}

// NewUserService is a function to initialize a UserService instance
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// All is a service for showing users
func (svc *UserService) All(params map[string]interface{}) (models.Users, error) {
	users, err := svc.repo.List(params)
	return users, err
}

// Show is a service for showing a user by ID
func (svc *UserService) Show(id int) (models.User, error) {
	user, err := svc.repo.Find(id)
	return user, err
}

// Create is a service for creating a new user
func (svc *UserService) Create(params inputs.UserCreateInput) (models.User, error) {
	user, err := svc.repo.Create(params.Output())
	return user, err
}

// Update is a service for updating an user
func (svc *UserService) Update(id int, params inputs.UserUpdateInput) (models.User, error) {
	user, err := svc.repo.Update(id, params.Output())
	return user, err
}

// Destroy is a service for destroying an user
func (svc *UserService) Destroy(id int) (models.User, error) {
	user, err := svc.repo.Destroy(id)
	return user, err
}
