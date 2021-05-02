package entity

import (
	"testing"

	"github.com/fikrimohammad/ficree-api/infrastructure/storage"
	"github.com/stretchr/testify/assert"
)

func TestAWSStoredFile(t *testing.T) {
	bucketName := storage.GetAWSInstance().Bucket.Name
	bucketURL := storage.GetAWSInstance().Bucket.URL
	tests := []struct {
		Name      string
		Input     string
		PublicURL string
	}{
		{
			Name:      "when the input is an AWS S3 Public URL and the subdomain is our bucket name",
			Input:     "https://" + bucketName + ".s3.ap-southeast-1.amazonaws.com/video/dummy_video.mp4",
			PublicURL: "https://" + bucketName + ".s3.ap-southeast-1.amazonaws.com/video/dummy_video.mp4",
		},
		{
			Name:      "when the input is a file path",
			Input:     "img nature.jpg",
			PublicURL: bucketURL + "/img%20nature.jpg",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(*testing.T) {
			file := NewAWSFile(test.Input)
			publicURL := file.PublicURL()
			uploadURL, _ := file.UploadURL()
			downloadURL, _ := file.DownloadURL()
			fileOutput, _ := file.AsFileOutput()

			assert.Equal(t, test.PublicURL, publicURL)

			assert.Contains(t, uploadURL, "X-Amz-SignedHeaders")
			assert.Contains(t, uploadURL, "X-Amz-Signature")
			assert.Contains(t, uploadURL, "X-Amz-Expires=900")
			assert.Contains(t, uploadURL, "X-Amz-Algorithm=AWS4-HMAC-SHA256")

			assert.Contains(t, downloadURL, "X-Amz-SignedHeaders")
			assert.Contains(t, downloadURL, "X-Amz-Signature")
			assert.Contains(t, downloadURL, "X-Amz-Expires=900")
			assert.Contains(t, downloadURL, "X-Amz-Algorithm=AWS4-HMAC-SHA256")

			assert.Equal(t, test.PublicURL, fileOutput.PublicURL)
		})
	}
}

func TestAWSUnstoredFile(t *testing.T) {
	unknownBucketName := "invalid_bucket"
	tests := []struct {
		Name      string
		Input     string
		PublicURL string
	}{
		{
			Name:      "when the input is an AWS S3 Public URL and the subdomain isn't our bucket name",
			Input:     "https://" + unknownBucketName + ".s3.ap-southeast-1.amazonaws.com/video/dummy_video.mp4",
			PublicURL: "https://" + unknownBucketName + ".s3.ap-southeast-1.amazonaws.com/video/dummy_video.mp4",
		},
		{
			Name:      "when the input is an URL but not a AWS S3 Public URL",
			Input:     "https://www.w3schools.com/w3css/img_nature.jpg",
			PublicURL: "https://www.w3schools.com/w3css/img_nature.jpg",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(*testing.T) {
			file := NewAWSFile(test.Input)
			publicURL := file.PublicURL()
			uploadURL, _ := file.UploadURL()
			downloadURL, _ := file.DownloadURL()
			fileOutput, _ := file.AsFileOutput()

			assert.Equal(t, test.PublicURL, publicURL)
			assert.Equal(t, test.PublicURL, downloadURL)
			assert.Equal(t, "", uploadURL)

			assert.Equal(t, test.PublicURL, fileOutput.PublicURL)
			assert.Equal(t, "", fileOutput.PresignedUploadURL)
		})
	}
}
