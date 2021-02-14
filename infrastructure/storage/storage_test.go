package storage

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/fikrimohammad/ficree-api/config"
	"github.com/stretchr/testify/assert"
)

func TestGetAWSInstance(t *testing.T) {
	instance := GetAWSInstance()
	assert.NotNil(t, instance)
}

func TestInitAWSInstance(t *testing.T) {
	t.Run("when force path style is true", func(*testing.T) {
		os.Setenv("AWS_FORCE_PATH_STYLE", "true")
		config.Load()

		InitAWSInstance()
		instance := GetAWSInstance()
		assert.NotNil(t, instance)
		assert.Equal(t, aws.Bool(true), instance.Config.S3ForcePathStyle)

		os.Unsetenv("AWS_FORCE_PATH_STYLE")
		config.Clear()
	})

	t.Run("when minio endpoint exists", func(t *testing.T) {
		minioEndpoint := "http://localhost:9000"

		os.Setenv("MINIO_ENDPOINT", minioEndpoint)
		config.Load()

		InitAWSInstance()
		instance := GetAWSInstance()
		assert.NotNil(t, instance)
		assert.Equal(t, aws.String(minioEndpoint), instance.Config.Endpoint)
		assert.Equal(t, aws.Bool(true), instance.Config.S3ForcePathStyle)

		os.Unsetenv("MINIO_ENDPOINT")
		config.Clear()
	})
}
