package presenters

import "github.com/fikrimohammad/ficree-api/app/models"

// EducationPresenter represents output builder for Education
type EducationPresenter struct {
	Education  models.Education
	FormatType string
}

// NewEducationPresenter is a function to initialize a EducationPresenter instance
func NewEducationPresenter(experience models.Education, formatType string) *EducationPresenter {
	presenter := EducationPresenter{Education: experience, FormatType: formatType}
	if presenter.FormatType == "" {
		presenter.FormatType = "format"
	}
	return &presenter
}

// Result is a function to select which format to be used for building the output
func (out *EducationPresenter) Result() map[string]interface{} {
	return out.format()
}

func (out *EducationPresenter) format() map[string]interface{} {
	output := map[string]interface{}{
		"id":                   out.Education.ID,
		"institution_name":     out.Education.InstitutionName,
		"institution_icon_url": out.Education.InstitutionIconURL,
		"institution_web_url":  out.Education.InstitutionWebURL,
		"description":          out.Education.Description,
		"starts_at":            out.Education.StartsAt,
		"ends_at":              out.Education.EndsAt,
		"degree":               out.Education.Degree,
		"study_field":          out.Education.StudyField,
	}
	return output
}
