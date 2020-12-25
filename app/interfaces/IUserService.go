package interfaces

import (
	"github.com/fikrimohammad/ficree-api/app/inputs"
	"github.com/fikrimohammad/ficree-api/app/models"
)

// IUserService represents UserService
type IUserService interface {
	All(params map[string]interface{}) (models.Users, error)
	Show(id int) (models.User, error)
	Create(params inputs.UserCreateInput) (models.User, error)
	Update(id int, params inputs.UserUpdateInput) (models.User, error)
	Destroy(id int) (models.User, error)
}
