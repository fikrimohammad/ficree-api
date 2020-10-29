package inputs

import (
	"time"

	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/gofiber/fiber/v2"
)

// ExperienceUpdateInput is a struct to store params for creating a new experience
type ExperienceUpdateInput struct {
	PositionName string    `json:"position_name"`
	StartsAt     time.Time `json:"starts_at"`
	EndsAt       time.Time `json:"ends_at"`
	Description  string    `json:"description"`
	CompanyName  string    `json:"company_name"`
	Location     string    `json:"location"`
}

// NewExperienceUpdateInput is a function to initialize a ExperienceUpdateInput instance
func NewExperienceUpdateInput(c *fiber.Ctx) (ExperienceUpdateInput, error) {
	input := ExperienceUpdateInput{}
	err := c.BodyParser(&input)
	return input, err
}

// Output is a function to convert ExperienceUpdateInput into a Experience instance
func (i *ExperienceUpdateInput) Output() models.Experience {
	output := models.Experience{
		PositionName: i.PositionName,
		Description:  i.Description,
		EndsAt:       i.EndsAt,
		StartsAt:     i.StartsAt,
		CompanyName:  i.CompanyName,
		Location:     i.Location,
	}
	return output
}
