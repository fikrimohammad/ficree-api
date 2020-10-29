package repositories

import (
	"github.com/fikrimohammad/ficree-api/app/models"
	"gorm.io/gorm"
)

// ExperienceRepository is a struct to wrap database transaction for Experience
type ExperienceRepository struct {
	db *gorm.DB
}

// NewExperienceRepository is a function to initialize a ExperienceRepository instance
func NewExperienceRepository(conn *gorm.DB) *ExperienceRepository {
	return &ExperienceRepository{db: conn}
}

// List is a function to fetch all active experiences
func (repo *ExperienceRepository) List() (models.Experiences, error) {
	experiences := models.Experiences{}
	err := repo.db.Find(&experiences).Error
	return experiences, err
}

// Find is a function to find a experience by ID
func (repo *ExperienceRepository) Find(id int) (models.Experience, error) {
	experience := models.Experience{}
	err := repo.db.First(&experience, id).Error
	return experience, err
}

// Create is a function to store a new experience
func (repo *ExperienceRepository) Create(newExperience models.Experience) (models.Experience, error) {
	err := repo.db.Create(&newExperience).Error
	return newExperience, err
}

// Update is a function to store a modified experience
func (repo *ExperienceRepository) Update(id int, modifiedExperience models.Experience) (models.Experience, error) {
	experience, fetchErr := repo.Find(id)
	if fetchErr != nil {
		return experience, fetchErr
	}

	updateErr := repo.db.Model(&experience).Updates(modifiedExperience).Error
	return experience, updateErr
}

// Destroy is a function to delete a experience
func (repo *ExperienceRepository) Destroy(id int) (models.Experience, error) {
	experience, fetchErr := repo.Find(id)
	if fetchErr != nil {
		return experience, fetchErr
	}

	deleteErr := repo.db.Delete(&experience).Error
	return experience, deleteErr
}
