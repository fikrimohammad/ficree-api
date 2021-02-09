package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserListInput_Validate(t *testing.T) {
	tests := []struct {
		Input       UserListInput
		Expectation map[string]interface{}
	}{
		{
			Input: UserListInput{
				Limit:         1,
				Offset:        1,
				SortColumn:    "name",
				SortDirection: "asc",
				SearchString:  "Soft",
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
		result := test.Input.AsQueryParams()
		for key, expectedValue := range test.Expectation {
			assert.Equal(t, expectedValue, result[key])
		}
	}
}
