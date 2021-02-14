package input

import "github.com/mitchellh/mapstructure"

// UserListInput represents supported query params for filtering users
type UserListInput struct {
	SearchString  string `mapstructure:"search_string,omitempty"`
	Limit         int    `mapstructure:"limit,omitempty"`
	Offset        int    `mapstructure:"offset,omitempty"`
	SortColumn    string `mapstructure:"sort_column,omitempty"`
	SortDirection string `mapstructure:"sort_direction,omitempty"`
}

// NewUserListInput is a function to initialize UserListInput instance
func NewUserListInput(params map[string]interface{}) (*UserListInput, error) {
	var input UserListInput
	err := mapstructure.Decode(params, &input)
	return &input, err
}

// AsQueryParams is a function to map parameters into query parameters
func (input *UserListInput) AsQueryParams() map[string]interface{} {
	return map[string]interface{}{
		"searchString":  input.SearchString,
		"limit":         input.Limit,
		"offset":        input.Offset,
		"sortColumn":    input.SortColumn,
		"sortDirection": input.SortDirection,
	}
}
