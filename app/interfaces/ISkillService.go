package interfaces

import (
	inputs "github.com/fikrimohammad/ficree-api/app/inputs/skills"
	"github.com/fikrimohammad/ficree-api/app/models"
)

// ISkillService represents SkillService
type ISkillService interface {
	All() (models.Skills, error)
	Show(id int) (models.Skill, error)
	Create(params inputs.SkillCreateInput) (models.Skill, error)
	Update(id int, params inputs.SkillUpdateInput) (models.Skill, error)
	Destroy(id int) (models.Skill, error)
}
