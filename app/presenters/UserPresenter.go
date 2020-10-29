package presenters

import (
	"github.com/fikrimohammad/ficree-api/app/models"
)

// UserPresenter represents output builder for User
type UserPresenter struct {
	User       models.User
	FormatType string
}

// NewUserPresenter is a function to initialize a UserPresenter instance
func NewUserPresenter(user models.User, formatType string) *UserPresenter {
	presenter := UserPresenter{User: user, FormatType: formatType}
	if presenter.FormatType == "" {
		presenter.FormatType = "format"
	}
	return &presenter
}

// Result is a function to select which format to be used for building the output
func (out *UserPresenter) Result() map[string]interface{} {
	if out.FormatType == "minimalFormat" {
		return out.minimalFormat()
	}
	return out.format()
}

func (out *UserPresenter) minimalFormat() map[string]interface{} {
	output := map[string]interface{}{
		"id":              out.User.ID,
		"name":            out.User.Name,
		"email":           out.User.Email,
		"title":           out.User.Title,
		"profile_picture": out.User.ProfilePicture,
	}
	return output
}

func (out *UserPresenter) format() map[string]interface{} {
	output := out.minimalFormat()
	output["github_url"] = out.User.GithubURL
	output["linkedin_url"] = out.User.LinkedinURL
	output["twitter_url"] = out.User.TwitterURL
	output["summary"] = out.User.Summary
	output["skills"] = out.skillOutput()
	output["experiences"] = out.experienceOutput()
	return output
}

func (out *UserPresenter) skillOutput() []map[string]interface{} {
	skills := out.User.Skills
	outputs := []map[string]interface{}{}
	for _, skill := range skills {
		output := NewSkillPresenter(skill, "").Result()
		outputs = append(outputs, output)
	}
	return outputs
}

func (out *UserPresenter) experienceOutput() []map[string]interface{} {
	experiences := out.User.Experiences
	outputs := []map[string]interface{}{}
	for _, experience := range experiences {
		output := NewExperiencePresenter(experience, "").Result()
		outputs = append(outputs, output)
	}
	return outputs
}
