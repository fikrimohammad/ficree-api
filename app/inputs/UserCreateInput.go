package inputs

import (
	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/gofiber/fiber/v2"
)

// UserCreateInput is a struct to store params for creating a new user
type UserCreateInput struct {
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

// NewUserCreateInput is a function to initialize a UserCreateInput instance
func NewUserCreateInput(c *fiber.Ctx) (UserCreateInput, error) {
	input := UserCreateInput{}
	err := c.BodyParser(&input)
	return input, err
}

// Output is a function to convert UserCreateInput into a User instance
func (i *UserCreateInput) Output() models.User {
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
