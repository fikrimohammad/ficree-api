package inputs

import (
	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/gofiber/fiber/v2"
)

// SkillUpdateInput is a struct to store params for creating a new skill
type SkillUpdateInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}

// NewSkillUpdateInput is a function to initialize a SkillUpdateInput instance
func NewSkillUpdateInput(c *fiber.Ctx) (SkillUpdateInput, error) {
	input := SkillUpdateInput{}
	err := c.BodyParser(&input)
	return input, err
}

// Output is a function to convert SkillUpdateInput into a Skill instance
func (i *SkillUpdateInput) Output() models.Skill {
	output := models.Skill{
		Name:        i.Name,
		Description: i.Description,
		Rating:      i.Rating,
	}
	return output
}
