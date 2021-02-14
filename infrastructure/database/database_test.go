package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	conn := Load()
	assert.NotNil(t, conn)

	Close()
}

func TestConnect(t *testing.T) {
	Connect()
	conn := Load()
	assert.NotNil(t, conn)

	Close()
}
