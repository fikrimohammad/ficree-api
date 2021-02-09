package domain

import (
	"time"

	"github.com/google/uuid"
)

// User represents user entity
type User struct {
	ID             int       `faker:"-"`
	GUID           uuid.UUID `faker:"-"`
	Name           string    `faker:"name"`
	Email          string    `faker:"email,unique"`
	PhoneNumber    string    `faker:"e_164_phone_number,unique"`
	ProfilePicture string    `faker:"url"`
	GithubURL      string    `faker:"url"`
	LinkedinURL    string    `faker:"url"`
	TwitterURL     string    `faker:"url"`
	Summary        string    `faker:"paragraph"`
	Title          string    `faker:"sentence"`
	CreatedAt      time.Time `pg:"created_at,default:now()" faker:"-"`
	UpdatedAt      time.Time `pg:"updated_at,default:now()" faker:"-"`
	DeletedAt      time.Time `pg:",soft_delete" faker:"-"`
}

// UserService represents usecase layer for processing users entity
type UserService interface {
	All(params map[string]interface{}) ([]map[string]interface{}, error)
	Show(id int) (map[string]interface{}, error)
	Create(params map[string]interface{}) (map[string]interface{}, error)
	Update(id int, params map[string]interface{}) (map[string]interface{}, error)
	Destroy(id int) (map[string]interface{}, error)
}

// UserRepository represents database layer for processing users entity
type UserRepository interface {
	List(params map[string]interface{}) ([]*User, error)
	Find(id int) (*User, error)
	Create(params *User) (*User, error)
	Update(id int, params *User) (*User, error)
	Destroy(id int) (*User, error)
}

// User error message template
const (
	FindUsersError    = "failed to find users: %v"
	FindUserByIDError = "failed to find user by ID: %v"
	CreateUserError   = "failed to create user: %v"
	UpdateUserError   = "failed to update user: %v"
	DeleteUserError   = "failed to delete user: %v"
)