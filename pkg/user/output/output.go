package output

import (
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/pkg/file/entity"
	"github.com/mitchellh/mapstructure"
)

// UserCompactOutput represents compact form of user output
type UserCompactOutput struct {
	ID             int    `mapstructure:"id"`
	Name           string `mapstructure:"name"`
	Email          string `mapstructure:"email"`
	Title          string `mapstructure:"title"`
	ProfilePicture string `mapstructure:"profile_picture"`
}

// UserDetailOutput represents detail form of user output
type UserDetailOutput struct {
	ID             int    `mapstructure:"id"`
	Name           string `mapstructure:"name"`
	Email          string `mapstructure:"email"`
	Title          string `mapstructure:"title"`
	ProfilePicture string `mapstructure:"profile_picture"`
	PhoneNumber    string `mapstructure:"phone_number"`
	GithubURL      string `mapstructure:"github_url"`
	LinkedinURL    string `mapstructure:"linkedin_url"`
	TwitterURL     string `mapstructure:"twitter_url"`
	Summary        string `mapstructure:"summary"`
}

// NewUserOutput is a function for creating user output
func NewUserOutput(user *domain.User, outputType string) (map[string]interface{}, error) {
	if outputType == "detail" {
		return asUserDetailOutput(user)
	}
	return asUserCompactOutput(user)
}

// NewUserArrayOutput is a function for creating array of user output
func NewUserArrayOutput(users []*domain.User, outputType string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	for _, user := range users {
		result, err := NewUserOutput(user, outputType)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func asUserCompactOutput(user *domain.User) (map[string]interface{}, error) {
	var result map[string]interface{}
	compactOutput := &UserCompactOutput{
		ID:             int(user.ID),
		Name:           user.Name,
		Email:          user.Email,
		Title:          user.Title,
		ProfilePicture: user.ProfilePicture,
	}

	profilePictureURL, err := buildPresignedDownloadURL(user.ProfilePicture)
	if err != nil {
		return nil, err
	}
	compactOutput.ProfilePicture = profilePictureURL

	err = mapstructure.Decode(compactOutput, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func asUserDetailOutput(user *domain.User) (map[string]interface{}, error) {
	var result map[string]interface{}
	detailOutput := &UserDetailOutput{
		ID:             int(user.ID),
		Name:           user.Name,
		Email:          user.Email,
		Title:          user.Title,
		Summary:        user.Summary,
		PhoneNumber:    user.PhoneNumber,
		GithubURL:      user.GithubURL,
		TwitterURL:     user.TwitterURL,
		LinkedinURL:    user.LinkedinURL,
		ProfilePicture: user.ProfilePicture,
	}

	profilePictureURL, err := buildPresignedDownloadURL(user.ProfilePicture)
	if err != nil {
		return nil, err
	}
	detailOutput.ProfilePicture = profilePictureURL

	err = mapstructure.Decode(detailOutput, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func buildPresignedDownloadURL(uri string) (string, error) {
	if uri == "" {
		return "", nil
	}

	object, err := entity.NewAWSFile(uri)
	if err != nil {
		return "", err
	}

	presignedURL, err := object.DownloadURL()
	if err != nil {
		return "", err
	}

	return presignedURL, nil
}
