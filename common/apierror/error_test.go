package apierror

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	errMsg := "dummy error"
	err := New(http.StatusInternalServerError, errMsg)
	assert.Error(t, err)
	assert.Equal(t, errMsg, err.Error())
}

func TestFromError(t *testing.T) {
	dummyErr := errors.New("dummy error")
	err := FromError(http.StatusInternalServerError, dummyErr)
	assert.Error(t, err)
	assert.Equal(t, dummyErr.Error(), err.Error())
}

func TestGetHTTPStatus(t *testing.T) {
	t.Run("when the error is an APIError instance", func(t *testing.T) {
		err := New(http.StatusBadRequest, "dummy error")
		assert.Error(t, err)

		statusCode := GetHTTPStatus(err)
		assert.Equal(t, http.StatusBadRequest, statusCode)
	})

	t.Run("when the error isn't an APIError instance", func(t *testing.T) {
		err := errors.New("dummy error")
		assert.Error(t, err)

		statusCode := GetHTTPStatus(err)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
	})
}
