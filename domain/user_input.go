package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// UserCreateInput is a struct to store params for creating a new user
type UserCreateInput struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	GithubURL      string `json:"github_url,omitempty"`
	LinkedinURL    string `json:"linkedin_url,omitempty"`
	TwitterURL     string `json:"twitter_url,omitempty"`
	Summary        string `json:"summary,omitempty"`
	Title          string `json:"title,omitempty"`
}

// Validate is a function to validate UserCreateInput values
func (i *UserCreateInput) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.Name, validation.Required, is.Alpha),
		validation.Field(&i.Email, validation.Required, is.EmailFormat),
		validation.Field(&i.PhoneNumber, is.Digit),
		validation.Field(&i.GithubURL, is.URL),
		validation.Field(&i.TwitterURL, is.URL),
		validation.Field(&i.LinkedinURL, is.URL),
	)
}

// AsUser is a function to convert UserCreateInput into a User instance
func (i *UserCreateInput) AsUser() *User {
	return &User{
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

// UserListInput represents supported query params for filtering users
type UserListInput struct {
	SearchString  string `query:"search_string,omitempty"`
	Limit         int    `query:"limit,omitempty"`
	Offset        int    `query:"offset,omitempty"`
	SortColumn    string `query:"sort_column"`
	SortDirection string `query:"sort_direction"`
}

// UserUpdateInput is a struct to store parameters for storing a modified user
type UserUpdateInput struct {
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	PhoneNumber    string `json:"phone_number,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	GithubURL      string `json:"github_url,omitempty"`
	LinkedinURL    string `json:"linkedin_url,omitempty"`
	TwitterURL     string `json:"twitter_url,omitempty"`
	Summary        string `json:"summary,omitempty"`
	Title          string `json:"title,omitempty"`
}

// Validate is a function to validate UserUpdateInput values
func (i *UserUpdateInput) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.Name, is.Alpha),
		validation.Field(&i.Email, is.EmailFormat),
		validation.Field(&i.PhoneNumber, is.Digit),
		validation.Field(&i.GithubURL, is.URL),
		validation.Field(&i.TwitterURL, is.URL),
		validation.Field(&i.LinkedinURL, is.URL),
	)
}

// AsUser is a function to convert UserUpdateInput into a User instance
func (i *UserUpdateInput) AsUser() *User {
	return &User{
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
