package services

import (
	inputs "github.com/fikrimohammad/ficree-api/app/inputs/skills"
	"github.com/fikrimohammad/ficree-api/app/interfaces"
	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/fikrimohammad/ficree-api/app/repositories"
)

// SkillService is a struct to wrap service for Skill
type SkillService struct {
	repo interfaces.ISkillRepository
}

// NewSkillService is a function to initialize a SkillService instance
func NewSkillService(repo *repositories.SkillRepository) *SkillService {
	return &SkillService{repo: repo}
}

// All is a service for showing all active users
func (svc *SkillService) All() (models.Skills, error) {
	users, err := svc.repo.List()
	return users, err
}

// Show is a service for showing a user by ID
func (svc *SkillService) Show(id int) (models.Skill, error) {
	user, err := svc.repo.Find(id)
	return user, err
}

// Create is a service for creating a new user
func (svc *SkillService) Create(params inputs.SkillCreateInput) (models.Skill, error) {
	user, err := svc.repo.Create(params.Output())
	return user, err
}

// Update is a service for updating an user
func (svc *SkillService) Update(id int, params inputs.SkillUpdateInput) (models.Skill, error) {
	user, err := svc.repo.Update(id, params.Output())
	return user, err
}

// Destroy is a service for destroying an user
func (svc *SkillService) Destroy(id int) (models.Skill, error) {
	user, err := svc.repo.Destroy(id)
	return user, err
}
