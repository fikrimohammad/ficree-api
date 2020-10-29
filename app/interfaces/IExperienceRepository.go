package interfaces

import "github.com/fikrimohammad/ficree-api/app/models"

// IExperienceRepository represents ExperienceRepository
type IExperienceRepository interface {
	List() (models.Experiences, error)
	Find(id int) (models.Experience, error)
	Create(params models.Experience) (models.Experience, error)
	Update(id int, params models.Experience) (models.Experience, error)
	Destroy(id int) (models.Experience, error)
}
