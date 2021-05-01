package service

import (
	"net/http"

	"github.com/fikrimohammad/ficree-api/common/apierror"
	"github.com/fikrimohammad/ficree-api/domain"
)

// UserService is a struct to wrap service for User
type UserService struct {
	Repo domain.UserRepository
}

// NewUserService is a function to initialize a UserService instance
func NewUserService(repo domain.UserRepository) domain.UserService {
	return &UserService{Repo: repo}
}

// All is a service for showing users
func (svc *UserService) List(params domain.UserListInput) ([]*domain.UserCompactOutput, error) {
	users, err := svc.Repo.List(params)
	if err != nil {
		return nil, err
	}

	results := make([]*domain.UserCompactOutput, 0)
	for _, user := range users {
		results = append(results, user.AsUserCompactOutput())
	}
	return results, err
}

// Show is a service for showing a user by ID
func (svc *UserService) Show(id int) (*domain.UserDetailOutput, error) {
	user, err := svc.Repo.Find(id)
	if err != nil {
		return nil, err
	}

	return user.AsUserDetailOutput(), err
}

// Create is a service for creating a new user
func (svc *UserService) Create(params domain.UserCreateInput) (*domain.UserDetailOutput, error) {
	err := params.Validate()
	if err != nil {
		return nil, apierror.FromError(http.StatusBadRequest, err)
	}

	user, err := svc.Repo.Create(params.AsUser())
	if err != nil {
		return nil, err
	}

	return user.AsUserDetailOutput(), err
}

// Update is a service for updating an user
func (svc *UserService) Update(id int, params domain.UserUpdateInput) (*domain.UserDetailOutput, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	user, err := svc.Repo.Update(id, params.AsUser())
	if err != nil {
		return nil, err
	}

	return user.AsUserDetailOutput(), err
}

// Destroy is a service for destroying an user
func (svc *UserService) Destroy(id int) (*domain.UserDetailOutput, error) {
	user, err := svc.Repo.Destroy(id)
	if err != nil {
		return nil, err
	}

	return user.AsUserDetailOutput(), err
}
