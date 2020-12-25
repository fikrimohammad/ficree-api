package inputs

import (
	"github.com/gofiber/fiber/v2"
)

// UserListInput is a struct to store query string for listing users
type UserListInput struct {
	Name  string `query:"name"`
	Limit int    `query:"limit"`
}

// NewUserListInput is a function to initialize a UserListInput instance
func NewUserListInput(c *fiber.Ctx) (UserListInput, error) {
	input := UserListInput{}
	err := c.QueryParser(&input)
	return input, err
}

// Output is a function to convert UserListInput into a map[string]interface{} instance
func (i *UserListInput) Output() map[string]interface{} {
	output := make(map[string]interface{})
	output["name"] = i.Name
	output["limit"] = i.Limit
	return output
}
