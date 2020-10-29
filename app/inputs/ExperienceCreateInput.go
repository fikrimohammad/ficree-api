package inputs

import (
	"time"

	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/gofiber/fiber/v2"
)

// ExperienceCreateInput is a struct to store params for creating a new experience
type ExperienceCreateInput struct {
	PositionName string    `json:"position_name"`
	StartsAt     time.Time `json:"starts_at"`
	EndsAt       time.Time `json:"ends_at"`
	Description  string    `json:"description"`
	CompanyName  string    `json:"company_name"`
	Location     string    `json:"location"`
	UserID       int       `json:"user_id"`
}

// NewExperienceCreateInput is a function to initialize a ExperienceCreateInput instance
func NewExperienceCreateInput(c *fiber.Ctx) (ExperienceCreateInput, error) {
	input := ExperienceCreateInput{}
	err := c.BodyParser(&input)
	return input, err
}

// Output is a function to convert ExperienceCreateInput into a Experience instance
func (i *ExperienceCreateInput) Output() models.Experience {
	output := models.Experience{
		PositionName: i.PositionName,
		Description:  i.Description,
		EndsAt:       i.EndsAt,
		StartsAt:     i.StartsAt,
		CompanyName:  i.CompanyName,
		Location:     i.Location,
		UserID:       i.UserID,
	}
	return output
}
