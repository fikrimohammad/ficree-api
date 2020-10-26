package repositories

import (
	"github.com/fikrimohammad/ficree-api/app/models"
	"gorm.io/gorm"
)

// UserRepository is a struct to wrap database transaction for User
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository is a function to initialize a UserRepository instance
func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{db: conn}
}

// List is a function to fetch all active users
func (repo *UserRepository) List() (models.Users, error) {
	users := models.Users{}
	err := repo.db.Find(&users).Error
	return users, err
}

// Find is a function to find an user by ID
func (repo *UserRepository) Find(id int) (models.User, error) {
	user := models.User{}
	err := repo.db.Preload("Skills").First(&user, id).Error
	return user, err
}

// Create is a function to store a new user
func (repo *UserRepository) Create(newUser models.User) (models.User, error) {
	err := repo.db.Create(&newUser).Error
	return newUser, err
}

// Update is a function to store a modified user
func (repo *UserRepository) Update(id int, modifiedUser models.User) (models.User, error) {
	user, fetchErr := repo.Find(id)
	if fetchErr != nil {
		return user, fetchErr
	}

	updateErr := repo.db.Model(&user).Updates(modifiedUser).Error
	return user, updateErr
}

// Destroy is a function to delete an user
func (repo *UserRepository) Destroy(id int) (models.User, error) {
	user, fetchErr := repo.Find(id)
	if fetchErr != nil {
		return user, fetchErr
	}

	deleteErr := repo.db.Delete(&user).Error
	return user, deleteErr
}
