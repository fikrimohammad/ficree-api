package interfaces

import "github.com/fikrimohammad/ficree-api/app/models"

// IEducationRepository represents EducationRepository
type IEducationRepository interface {
	List() (models.Educations, error)
	Find(id int) (models.Education, error)
	Create(params models.Education) (models.Education, error)
	Update(id int, params models.Education) (models.Education, error)
	Destroy(id int) (models.Education, error)
}
