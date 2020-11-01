package repositories

import (
	"github.com/fikrimohammad/ficree-api/app/models"
	"gorm.io/gorm"
)

// EducationRepository is a struct to wrap database transaction for Education
type EducationRepository struct {
	db *gorm.DB
}

// NewEducationRepository is a function to initialize a EducationRepository instance
func NewEducationRepository(conn *gorm.DB) *EducationRepository {
	return &EducationRepository{db: conn}
}

// List is a function to fetch all active educationss
func (repo *EducationRepository) List() (models.Educations, error) {
	educationss := models.Educations{}
	err := repo.db.Find(&educationss).Error
	return educationss, err
}

// Find is a function to find a educations by ID
func (repo *EducationRepository) Find(id int) (models.Education, error) {
	educations := models.Education{}
	err := repo.db.First(&educations, id).Error
	return educations, err
}

// Create is a function to store a new educations
func (repo *EducationRepository) Create(newEducation models.Education) (models.Education, error) {
	err := repo.db.Create(&newEducation).Error
	return newEducation, err
}

// Update is a function to store a modified educations
func (repo *EducationRepository) Update(id int, modifiedEducation models.Education) (models.Education, error) {
	educations, fetchErr := repo.Find(id)
	if fetchErr != nil {
		return educations, fetchErr
	}

	updateErr := repo.db.Model(&educations).Updates(modifiedEducation).Error
	return educations, updateErr
}

// Destroy is a function to delete a educations
func (repo *EducationRepository) Destroy(id int) (models.Education, error) {
	educations, fetchErr := repo.Find(id)
	if fetchErr != nil {
		return educations, fetchErr
	}

	deleteErr := repo.db.Delete(&educations).Error
	return educations, deleteErr
}
