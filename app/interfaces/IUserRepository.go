package interfaces

import "github.com/fikrimohammad/ficree-api/app/models"

// IUserRepository represents UserRepository
type IUserRepository interface {
	List() (models.Users, error)
	Find(id int) (models.User, error)
	Create(params models.User) (models.User, error)
	Update(id int, params models.User) (models.User, error)
	Destroy(id int) (models.User, error)
}
