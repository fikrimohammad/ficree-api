package presenters

import "github.com/fikrimohammad/ficree-api/app/models"

// UserPresenter represents output builder for User
type UserPresenter struct {
	User       models.User
	FormatType string
}

// NewUserPresenter is a function to initialize a UserPresenter instance
func NewUserPresenter(user models.User, formatType string) *UserPresenter {
	return &UserPresenter{
		User:       user,
		FormatType: formatType,
	}
}

// Result is a function to select which format to be used for building the output
func (out *UserPresenter) Result() map[string]interface{} {
	if out.FormatType == "minimal_format" {
		return out.minimalFormat()
	}
	return out.detailFormat()
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

func (out *UserPresenter) detailFormat() map[string]interface{} {
	output := out.minimalFormat()
	output["github_url"] = out.User.GithubURL
	output["linkedin_url"] = out.User.LinkedinURL
	output["twitter_url"] = out.User.TwitterURL
	output["summary"] = out.User.Summary
	output["skills"] = out.skillOutput()
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
