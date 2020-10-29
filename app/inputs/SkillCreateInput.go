package inputs

import (
	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/gofiber/fiber/v2"
)

// SkillCreateInput is a struct to store params for creating a new skill
type SkillCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	UserID      int    `json:"user_id"`
}

// NewSkillCreateInput is a function to initialize a SkillCreateInput instance
func NewSkillCreateInput(c *fiber.Ctx) (SkillCreateInput, error) {
	input := SkillCreateInput{}
	err := c.BodyParser(&input)
	return input, err
}

// Output is a function to convert SkillCreateInput into a Skill instance
func (i *SkillCreateInput) Output() models.Skill {
	output := models.Skill{
		Name:        i.Name,
		Description: i.Description,
		Rating:      i.Rating,
		UserID:      i.UserID,
	}
	return output
}
