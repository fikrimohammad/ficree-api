package services

import (
	"github.com/fikrimohammad/ficree-api/app/inputs"
	"github.com/fikrimohammad/ficree-api/app/interfaces"
	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/fikrimohammad/ficree-api/app/repositories"
)

// ExperienceService is a struct to wrap service for Experience
type ExperienceService struct {
	repo interfaces.IExperienceRepository
}

// NewExperienceService is a function to initialize a ExperienceService instance
func NewExperienceService(repo *repositories.ExperienceRepository) *ExperienceService {
	return &ExperienceService{repo: repo}
}

// All is a service for showing all active experiences
func (svc *ExperienceService) All() (models.Experiences, error) {
	experiences, err := svc.repo.List()
	return experiences, err
}

// Show is a service for showing a experience by ID
func (svc *ExperienceService) Show(id int) (models.Experience, error) {
	experience, err := svc.repo.Find(id)
	return experience, err
}

// Create is a service for creating a new experience
func (svc *ExperienceService) Create(params inputs.ExperienceCreateInput) (models.Experience, error) {
	experience, err := svc.repo.Create(params.Output())
	return experience, err
}

// Update is a service for updating an experience
func (svc *ExperienceService) Update(id int, params inputs.ExperienceUpdateInput) (models.Experience, error) {
	experience, err := svc.repo.Update(id, params.Output())
	return experience, err
}

// Destroy is a service for destroying an experience
func (svc *ExperienceService) Destroy(id int) (models.Experience, error) {
	experience, err := svc.repo.Destroy(id)
	return experience, err
}
