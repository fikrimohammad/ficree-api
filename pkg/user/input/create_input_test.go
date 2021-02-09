package input

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserCreateInputSuite struct {
	suite.Suite
}

func TestUserCreateInputSuite(t *testing.T) {
	suite.Run(t, new(UserCreateInputSuite))
}

func (suite *UserCreateInputSuite) TestUserCreateInput_Validate() {
	suite.Run("when name input is empty", func() {
		params := map[string]interface{}{
			"name":         nil,
			"email":        "ficree@gmail.com",
			"phone_number": "62811111111",
		}

		input, err := NewUserCreateInput(params)
		suite.NoError(err)

		err = input.Validate()
		suite.Error(err)
	})

	suite.Run("when name input constains character other than alphabet", func() {
		params := map[string]interface{}{
			"name":         "Ficr33",
			"email":        "ficree@gmail.com",
			"phone_number": "62811111111",
		}

		input, err := NewUserCreateInput(params)
		suite.NoError(err)

		err = input.Validate()
		suite.Error(err)
	})

	suite.Run("when email input is empty", func() {
		params := map[string]interface{}{
			"name":         "Ficree",
			"email":        nil,
			"phone_number": "62811111111",
		}

		input, err := NewUserCreateInput(params)
		suite.NoError(err)

		err = input.Validate()
		suite.Error(err)
	})

	suite.Run("when email input isn't in right format", func() {
		params := map[string]interface{}{
			"name":         "Ficree",
			"email":        "invalid_email",
			"phone_number": "62811111111",
		}

		input, err := NewUserCreateInput(params)
		suite.NoError(err)

		err = input.Validate()
		suite.Error(err)
	})

	suite.Run("when phone number input constains character other than number", func() {
		params := map[string]interface{}{
			"name":         "Ficree",
			"email":        "ficree@gmail.com",
			"phone_number": "628IIIIIIIII",
		}

		input, err := NewUserCreateInput(params)
		suite.NoError(err)

		err = input.Validate()
		suite.Error(err)
	})

	suite.Run("when github URL input isn't an URL", func() {
		params := map[string]interface{}{
			"name":         "Ficree",
			"email":        "ficree@gmail.com",
			"phone_number": "628111111111",
			"github_url":   "invalid_url",
		}

		input, err := NewUserCreateInput(params)
		suite.NoError(err)

		err = input.Validate()
		suite.Error(err)
	})

	suite.Run("when twitter URL input isn't an URL", func() {
		params := map[string]interface{}{
			"name":         "Ficree",
			"email":        "ficree@gmail.com",
			"phone_number": "628111111111",
			"twitter_url":  "invalid_url",
		}

		input, err := NewUserCreateInput(params)
		suite.NoError(err)

		err = input.Validate()
		suite.Error(err)
	})

	suite.Run("when linkedin URL input isn't an URL", func() {
		params := map[string]interface{}{
			"name":         "Ficree",
			"email":        "ficree@gmail.com",
			"phone_number": "628111111111",
			"linkedin_url": "invalid_url",
		}

		input, err := NewUserCreateInput(params)
		suite.NoError(err)

		err = input.Validate()
		suite.Error(err)
	})
}

func (suite *UserCreateInputSuite) TestUserCreateInput_AsUser() {
	params := map[string]interface{}{
		"name":         "Ficree",
		"email":        "ficree@gmail.com",
		"phone_number": "628111111111",
	}

	input, err := NewUserCreateInput(params)
	suite.NoError(err)

	user := input.AsUser()
	suite.Equal(input.Name, user.Name)
	suite.Equal(input.Email, user.Email)
	suite.Equal(input.PhoneNumber, user.PhoneNumber)
}
