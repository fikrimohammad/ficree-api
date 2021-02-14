package entity

import (
	"testing"

	"github.com/fikrimohammad/ficree-api/infrastructure/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewAWSFile(t *testing.T) {
	bucketName := storage.GetAWSInstance().Bucket.Name
	unknownBucketName := "unknown_bucket"

	tests := []struct {
		Name       string
		Input      string
		ResultType string
	}{
		{
			Name:       "when the input is an AWS S3 Public URL and the subdomain is our bucket name",
			Input:      "https://" + bucketName + ".s3.ap-southeast-1.amazonaws.com/video/dummy_video.mp4",
			ResultType: "aws-stored",
		},
		{
			Name:       "when the input is a file path",
			Input:      "img nature.jpg",
			ResultType: "aws-stored",
		},
		{
			Name:       "when the input is an AWS S3 Public URL and the subdomain isn't our bucket name",
			Input:      "https://" + unknownBucketName + ".s3.ap-southeast-1.amazonaws.com/video/dummy_video.mp4",
			ResultType: "unstored",
		},
		{
			Name:       "when the input is an URL but not a AWS S3 Public URL",
			Input:      "https://www.w3schools.com/w3css/img_nature.jpg",
			ResultType: "unstored",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(*testing.T) {
			result, err := NewAWSFile(test.Input)

			switch test.ResultType {
			case "aws-stored":
				assert.NoError(t, err)
				assert.IsType(t, &AWSStoredFile{}, result)
			case "unstored":
				assert.NoError(t, err)
				assert.IsType(t, &AWSUnStoredFile{}, result)
			default:
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}
