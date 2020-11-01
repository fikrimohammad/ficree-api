package interfaces

import (
	"github.com/fikrimohammad/ficree-api/app/inputs"
	"github.com/fikrimohammad/ficree-api/app/models"
)

// IEducationService represents EducationService
type IEducationService interface {
	All() (models.Educations, error)
	Show(id int) (models.Education, error)
	Create(params inputs.EducationCreateInput) (models.Education, error)
	Update(id int, params inputs.EducationUpdateInput) (models.Education, error)
	Destroy(id int) (models.Education, error)
}
