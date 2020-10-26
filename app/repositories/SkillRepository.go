package repositories

import (
	"github.com/fikrimohammad/ficree-api/app/models"
	"gorm.io/gorm"
)

// SkillRepository is a struct to wrap database transaction for Skill
type SkillRepository struct {
	db *gorm.DB
}

// NewSkillRepository is a function to initialize a SkillRepository instance
func NewSkillRepository(conn *gorm.DB) *SkillRepository {
	return &SkillRepository{db: conn}
}

// List is a function to fetch all active skills
func (repo *SkillRepository) List() (models.Skills, error) {
	skills := models.Skills{}
	err := repo.db.Find(&skills).Error
	return skills, err
}

// Find is a function to find a skill by ID
func (repo *SkillRepository) Find(id int) (models.Skill, error) {
	skill := models.Skill{}
	err := repo.db.First(&skill, id).Error
	return skill, err
}

// Create is a function to store a new skill
func (repo *SkillRepository) Create(newSkill models.Skill) (models.Skill, error) {
	err := repo.db.Create(&newSkill).Error
	return newSkill, err
}

// Update is a function to store a modified skill
func (repo *SkillRepository) Update(id int, modifiedSkill models.Skill) (models.Skill, error) {
	skill, fetchErr := repo.Find(id)
	if fetchErr != nil {
		return skill, fetchErr
	}

	updateErr := repo.db.Model(&skill).Updates(modifiedSkill).Error
	return skill, updateErr
}

// Destroy is a function to delete a skill
func (repo *SkillRepository) Destroy(id int) (models.Skill, error) {
	skill, fetchErr := repo.Find(id)
	if fetchErr != nil {
		return skill, fetchErr
	}

	deleteErr := repo.db.Delete(&skill).Error
	return skill, deleteErr
}
