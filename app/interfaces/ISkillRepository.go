package interfaces

import "github.com/fikrimohammad/ficree-api/app/models"

// ISkillRepository represents SkillRepository
type ISkillRepository interface {
	List() (models.Skills, error)
	Find(id int) (models.Skill, error)
	Create(params models.Skill) (models.Skill, error)
	Update(id int, params models.Skill) (models.Skill, error)
	Destroy(id int) (models.Skill, error)
}
