package input

import (
	"github.com/fikrimohammad/ficree-api/domain"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mitchellh/mapstructure"
)

// UserUpdateInput is a struct to store parameters for storing a modified user
type UserUpdateInput struct {
	Name           string `mapstructure:"name,omitempty"`
	Email          string `mapstructure:"email,omitempty"`
	PhoneNumber    string `mapstructure:"phone_number,omitempty"`
	ProfilePicture string `mapstructure:"profile_picture,omitempty"`
	GithubURL      string `mapstructure:"github_url,omitempty"`
	LinkedinURL    string `mapstructure:"linkedin_url,omitempty"`
	TwitterURL     string `mapstructure:"twitter_url,omitempty"`
	Summary        string `mapstructure:"summary,omitempty"`
	Title          string `mapstructure:"title,omitempty"`
}

// NewUserUpdateInput is a function to initialize a UserUpdateInput instance
func NewUserUpdateInput(params map[string]interface{}) (*UserUpdateInput, error) {
	var input UserUpdateInput
	err := mapstructure.Decode(params, &input)
	return &input, err
}

// Validate is a function to validate UserUpdateInput values
func (i UserUpdateInput) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Name, is.Alpha),
		validation.Field(&i.Email, is.EmailFormat),
		validation.Field(&i.PhoneNumber, is.Digit),
		validation.Field(&i.GithubURL, is.URL),
		validation.Field(&i.TwitterURL, is.URL),
		validation.Field(&i.LinkedinURL, is.URL),
	)
}

// AsUser is a function to convert UserUpdateInput into a User instance
func (i *UserUpdateInput) AsUser() *domain.User {
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
