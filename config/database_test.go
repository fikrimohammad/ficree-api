package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDBConfig(t *testing.T) {
	os.Setenv("TEST_DB_HOST", "192.168.1.12")
	os.Setenv("TEST_DB_PORT", "5431")
	os.Setenv("TEST_DB_USER", "ficree_test")
	os.Setenv("TEST_DB_PASSWORD", "fortesting")
	os.Setenv("TEST_DB_NAME", "ficree_test")

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "ficree_dev")
	os.Setenv("DB_PASSWORD", "fordevelopment")
	os.Setenv("DB_NAME", "ficree_development")

	t.Run("when app env is testing", func(*testing.T) {
		os.Setenv("APP_ENV", "test")
		Load()

		assert.Equal(t, "192.168.1.12", GetDBHost())
		assert.Equal(t, "5431", GetDBPort())
		assert.Equal(t, "ficree_test", GetDBUsername())
		assert.Equal(t, "fortesting", GetDBPassword())
		assert.Equal(t, "ficree_test", GetDBName())
	})

	t.Run("when app env is dev or production", func(t *testing.T) {
		os.Setenv("APP_ENV", "development")
		Load()

		assert.Equal(t, "localhost", GetDBHost())
		assert.Equal(t, "5432", GetDBPort())
		assert.Equal(t, "ficree_dev", GetDBUsername())
		assert.Equal(t, "fordevelopment", GetDBPassword())
		assert.Equal(t, "ficree_development", GetDBName())
	})

	os.Unsetenv("TEST_DB_HOST")
	os.Unsetenv("TEST_DB_PORT")
	os.Unsetenv("TEST_DB_USER")
	os.Unsetenv("TEST_DB_PASSWORD")
	os.Unsetenv("TEST_DB_NAME")

	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")

	Clear()
}
