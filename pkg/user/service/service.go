package service

import (
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/pkg/user/input"
	"github.com/fikrimohammad/ficree-api/pkg/user/output"
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
func (svc *UserService) All(params map[string]interface{}) ([]map[string]interface{}, error) {
	listInput, err := input.NewUserListInput(params)
	if err != nil {
		return nil, err
	}

	users, err := svc.Repo.List(listInput.AsQueryParams())
	if err != nil {
		return nil, err
	}

	results, err := output.NewUserArrayOutput(users, domain.UserCompactOutputType)
	return results, err
}

// Show is a service for showing a user by ID
func (svc *UserService) Show(id int) (map[string]interface{}, error) {
	user, err := svc.Repo.Find(id)
	if err != nil {
		return nil, err
	}

	result, err := output.NewUserOutput(user, domain.UserDetailOutputType)
	return result, err
}

// Create is a service for creating a new user
func (svc *UserService) Create(params map[string]interface{}) (map[string]interface{}, error) {
	createInput, err := input.NewUserCreateInput(params)
	if err != nil {
		return nil, err
	}

	err = createInput.Validate()
	if err != nil {
		return nil, err
	}

	user, err := svc.Repo.Create(createInput.AsUser())
	if err != nil {
		return nil, err
	}

	userOut, err := output.NewUserOutput(user, domain.UserCompactOutputType)
	return userOut, err
}

// Update is a service for updating an user
func (svc *UserService) Update(id int, params map[string]interface{}) (map[string]interface{}, error) {
	updateInput, err := input.NewUserUpdateInput(params)
	if err != nil {
		return nil, err
	}

	err = updateInput.Validate()
	if err != nil {
		return nil, err
	}

	user, err := svc.Repo.Update(id, updateInput.AsUser())
	if err != nil {
		return nil, err
	}

	userOut, err := output.NewUserOutput(user, domain.UserCompactOutputType)
	return userOut, err
}

// Destroy is a service for destroying an user
func (svc *UserService) Destroy(id int) (map[string]interface{}, error) {
	user, err := svc.Repo.Destroy(id)
	if err != nil {
		return nil, err
	}

	userOut, err := output.NewUserOutput(user, domain.UserCompactOutputType)
	return userOut, err
}
