package interfaces

import "github.com/fikrimohammad/ficree-api/app/models"

type IUserRepository interface {
	Find(id int) (models.User, error)
	Create(params models.User) (models.User, error)
	Update(id int, params models.User) (models.User, error)
	Destroy(id int) (models.User, error)
}
