package input

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserUpdateInputSuite struct {
	suite.Suite
}

func TestUserUpdateInputSuite(t *testing.T) {
	suite.Run(t, new(UserUpdateInputSuite))
}

func (suite *UserUpdateInputSuite) TestUserUpdateInput_Validate() {
	tests := []struct {
		Name    string
		Params  map[string]interface{}
		IsError bool
	}{
		{
			Name: "when name input constains character other than alphabet",
			Params: map[string]interface{}{
				"name": "Ficr33",
			},
			IsError: true,
		},
		{
			Name: "when email input isn't in right format",
			Params: map[string]interface{}{
				"email": "invalid email",
			},
			IsError: true,
		},
		{
			Name: "when phone number input constains character other than number",
			Params: map[string]interface{}{
				"phone_number": "628IIIIIIIII",
			},
			IsError: true,
		},
		{
			Name: "when github URL input isn't an URL",
			Params: map[string]interface{}{
				"github_url": "invalid url",
			},
			IsError: true,
		},
		{
			Name: "when twitter URL input isn't an URL",
			Params: map[string]interface{}{
				"twitter_url": "invalid url",
			},
			IsError: true,
		},
		{
			Name: "when linkedin URL input isn't an URL",
			Params: map[string]interface{}{
				"linkedin_url": "invalid url",
			},
			IsError: true,
		},
	}

	for _, test := range tests {
		suite.Run(test.Name, func() {
			input, err := NewUserUpdateInput(test.Params)
			suite.NoError(err)

			err = input.Validate()
			if test.IsError {
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
		})
	}
}

func (suite *UserUpdateInputSuite) TestUserUpdateInput_AsUser() {
	params := map[string]interface{}{
		"name":         "Ficree",
		"email":        "ficree@gmail.com",
		"phone_number": "628111111111",
	}

	input, err := NewUserUpdateInput(params)
	suite.NoError(err)

	user := input.AsUser()
	suite.Equal(input.Name, user.Name)
	suite.Equal(input.Email, user.Email)
	suite.Equal(input.PhoneNumber, user.PhoneNumber)
}
