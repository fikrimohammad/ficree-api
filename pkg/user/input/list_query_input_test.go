package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserListInput_AsQueryParams(t *testing.T) {
	tests := []struct {
		Input       map[string]interface{}
		Expectation map[string]interface{}
	}{
		{
			Input: map[string]interface{}{
				"search_string":  "Soft",
				"limit":          1,
				"offset":         1,
				"sort_column":    "name",
				"sort_direction": "asc",
			},
			Expectation: map[string]interface{}{
				"searchString":  "Soft",
				"limit":         1,
				"offset":        1,
				"sortColumn":    "name",
				"sortDirection": "asc",
			},
		},
	}

	for _, test := range tests {
		input, err := NewUserListInput(test.Input)
		assert.NoError(t, err)

		result := input.AsQueryParams()
		for key, expectedValue := range test.Expectation {
			assert.Equal(t, expectedValue, result[key])
		}
	}
}
