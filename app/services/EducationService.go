package services

import (
	"github.com/fikrimohammad/ficree-api/app/inputs"
	"github.com/fikrimohammad/ficree-api/app/interfaces"
	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/fikrimohammad/ficree-api/app/repositories"
)

// EducationService is a struct to wrap service for Education
type EducationService struct {
	repo interfaces.IEducationRepository
}

// NewEducationService is a function to initialize a EducationService instance
func NewEducationService(repo *repositories.EducationRepository) *EducationService {
	return &EducationService{repo: repo}
}

// All is a service for showing all active experiences
func (svc *EducationService) All() (models.Educations, error) {
	experiences, err := svc.repo.List()
	return experiences, err
}

// Show is a service for showing a experience by ID
func (svc *EducationService) Show(id int) (models.Education, error) {
	experience, err := svc.repo.Find(id)
	return experience, err
}

// Create is a service for creating a new experience
func (svc *EducationService) Create(params inputs.EducationCreateInput) (models.Education, error) {
	experience, err := svc.repo.Create(params.Output())
	return experience, err
}

// Update is a service for updating an experience
func (svc *EducationService) Update(id int, params inputs.EducationUpdateInput) (models.Education, error) {
	experience, err := svc.repo.Update(id, params.Output())
	return experience, err
}

// Destroy is a service for destroying an experience
func (svc *EducationService) Destroy(id int) (models.Education, error) {
	experience, err := svc.repo.Destroy(id)
	return experience, err
}
