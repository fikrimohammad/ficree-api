package presenters

import "github.com/fikrimohammad/ficree-api/app/models"

// SkillPresenter represents output builder for Skill
type SkillPresenter struct {
	Skill      models.Skill
	FormatType string
}

// NewSkillPresenter is a function to initialize a SkillPresenter instance
func NewSkillPresenter(skill models.Skill, formatType string) *SkillPresenter {
	return &SkillPresenter{
		Skill:      skill,
		FormatType: formatType,
	}
}

// Result is a function to select which format to be used for building the output
func (out *SkillPresenter) Result() map[string]interface{} {
	return out.format()
}

func (out *SkillPresenter) format() map[string]interface{} {
	output := map[string]interface{}{
		"id":          out.Skill.ID,
		"name":        out.Skill.Name,
		"rating":      out.Skill.Rating,
		"description": out.Skill.Description,
	}
	return output
}
