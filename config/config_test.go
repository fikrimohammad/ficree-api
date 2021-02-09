package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	cfg := Get()
	assert.NotNil(t, cfg)

	Clear()
}

func TestIsTesting(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	Load()

	isTesting := IsTesting()
	assert.Equal(t, true, isTesting)

	Clear()
}

func TestIsDevelopment(t *testing.T) {
	os.Setenv("APP_ENV", "development")
	Load()

	isDevelopment := IsDevelopment()
	assert.Equal(t, true, isDevelopment)

	Clear()
}

func TestIsProduction(t *testing.T) {
	os.Setenv("APP_ENV", "production")
	Load()

	isProduction := IsProduction()
	assert.Equal(t, true, isProduction)

	Clear()
}
