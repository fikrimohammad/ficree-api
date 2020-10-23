package inputs

import (
	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/gofiber/fiber/v2"
)

// UserUpdateInput is a struct to store parameters for storing a modified user
type UserUpdateInput struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	ProfilePicture string `json:"profile_picture"`
	GithubURL      string `json:"github_url"`
	LinkedinURL    string `json:"linkedin_url"`
	TwitterURL     string `json:"twitter_url"`
	Summary        string `json:"summary"`
	Title          string `json:"title"`
}

// NewUserUpdateInput is a function to initialize a UserUpdateInput instance
func NewUserUpdateInput(c *fiber.Ctx) (UserUpdateInput, error) {
	input := UserUpdateInput{}
	err := c.BodyParser(&input)
	return input, err
}

// Output is a function to convert UserUpdateInput into a User instance
func (i *UserUpdateInput) Output() models.User {
	output := models.User{
		Name:           i.Name,
		Email:          i.Email,
		PhoneNumber:    i.PhoneNumber,
		ProfilePicture: i.ProfilePicture,
		GithubURL:      i.GithubURL,
		TwitterURL:     i.TwitterURL,
		LinkedinURL:    i.LinkedinURL,
		Summary:        i.Summary,
		Title:          i.Title,
	}
	return output
}
