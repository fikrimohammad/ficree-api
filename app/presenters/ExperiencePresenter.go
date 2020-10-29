package presenters

import "github.com/fikrimohammad/ficree-api/app/models"

// ExperiencePresenter represents output builder for Experience
type ExperiencePresenter struct {
	Experience models.Experience
	FormatType string
}

// NewExperiencePresenter is a function to initialize a ExperiencePresenter instance
func NewExperiencePresenter(experience models.Experience, formatType string) *ExperiencePresenter {
	presenter := ExperiencePresenter{Experience: experience, FormatType: formatType}
	if presenter.FormatType == "" {
		presenter.FormatType = "format"
	}
	return &presenter
}

// Result is a function to select which format to be used for building the output
func (out *ExperiencePresenter) Result() map[string]interface{} {
	return out.format()
}

func (out *ExperiencePresenter) format() map[string]interface{} {
	output := map[string]interface{}{
		"id":            out.Experience.ID,
		"position_name": out.Experience.PositionName,
		"description":   out.Experience.Description,
		"starts_at":     out.Experience.StartsAt,
		"ends_at":       out.Experience.EndsAt,
		"company_name":  out.Experience.CompanyName,
		"location":      out.Experience.Location,
	}
	return output
}
