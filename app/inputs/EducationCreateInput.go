package inputs

import (
	"time"

	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/gofiber/fiber/v2"
)

// EducationCreateInput is a struct to store params for creating a new experience
type EducationCreateInput struct {
	InstitutionName    string    `json:"institution_name"`
	InstitutionIconURL string    `json:"institution_icon_url"`
	InstitutionWebURL  string    `json:"institution_web_url"`
	Description        string    `json:"description"`
	StartsAt           time.Time `json:"starts_at"`
	EndsAt             time.Time `json:"ends_at"`
	Degree             string    `json:"degree"`
	StudyField         string    `json:"study_field"`
	UserID             int       `json:"user_id"`
}

// NewEducationCreateInput is a function to initialize a EducationCreateInput instance
func NewEducationCreateInput(c *fiber.Ctx) (EducationCreateInput, error) {
	input := EducationCreateInput{}
	err := c.BodyParser(&input)
	return input, err
}

// Output is a function to convert EducationCreateInput into a Education instance
func (i *EducationCreateInput) Output() models.Education {
	output := models.Education{
		InstitutionName:    i.InstitutionName,
		InstitutionIconURL: i.InstitutionIconURL,
		InstitutionWebURL:  i.InstitutionWebURL,
		Description:        i.Description,
		EndsAt:             i.EndsAt,
		StartsAt:           i.StartsAt,
		Degree:             i.Degree,
		StudyField:         i.StudyField,
		UserID:             i.UserID,
	}
	return output
}
