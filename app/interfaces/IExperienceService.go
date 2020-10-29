package interfaces

import (
	"github.com/fikrimohammad/ficree-api/app/inputs"
	"github.com/fikrimohammad/ficree-api/app/models"
)

// IExperienceService represents ExperienceService
type IExperienceService interface {
	All() (models.Experiences, error)
	Show(id int) (models.Experience, error)
	Create(params inputs.ExperienceCreateInput) (models.Experience, error)
	Update(id int, params inputs.ExperienceUpdateInput) (models.Experience, error)
	Destroy(id int) (models.Experience, error)
}
