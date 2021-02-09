package input

import (
	"github.com/fikrimohammad/ficree-api/domain"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mitchellh/mapstructure"
)

// UserCreateInput is a struct to store params for creating a new user
type UserCreateInput struct {
	Name           string `mapstructure:"name"`
	Email          string `mapstructure:"email"`
	PhoneNumber    string `mapstructure:"phone_number"`
	ProfilePicture string `mapstructure:"profile_picture,omitempty"`
	GithubURL      string `mapstructure:"github_url,omitempty"`
	LinkedinURL    string `mapstructure:"linkedin_url,omitempty"`
	TwitterURL     string `mapstructure:"twitter_url,omitempty"`
	Summary        string `mapstructure:"summary,omitempty"`
	Title          string `mapstructure:"title,omitempty"`
}

// NewUserCreateInput is a function to initialize a UserCreateInput instance
func NewUserCreateInput(params map[string]interface{}) (*UserCreateInput, error) {
	var input UserCreateInput
	err := mapstructure.Decode(params, &input)
	return &input, err
}

// Validate is a function to validate UserCreateInput values
func (i UserCreateInput) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Name, validation.Required, is.Alpha),
		validation.Field(&i.Email, validation.Required, is.EmailFormat),
		validation.Field(&i.PhoneNumber, is.Digit),
		validation.Field(&i.GithubURL, is.URL),
		validation.Field(&i.TwitterURL, is.URL),
		validation.Field(&i.LinkedinURL, is.URL),
	)
}

// AsUser is a function to convert UserCreateInput into a User instance
func (i *UserCreateInput) AsUser() *domain.User {
	return &domain.User{
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
}
