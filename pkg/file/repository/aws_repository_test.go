package repository

import (
	"testing"

	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/infrastructure/storage"
	"github.com/stretchr/testify/suite"
)

type AWSFileRepositorySuite struct {
	suite.Suite
	repo domain.FileRepository
}

func (suite *AWSFileRepositorySuite) SetupSuite() {
	suite.repo = NewAWSFileRepository()
}

func TestAWSFileRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip tests for AWSFileRepositorySuite")
	}
	suite.Run(t, new(AWSFileRepositorySuite))
}

func (suite *AWSFileRepositorySuite) TestAWSFileRepositorySuite_FindByURI() {
	bucketName := storage.GetAWSInstance().Bucket.Name
	uri := "https://" + bucketName + ".s3.ap-southeast-1.amazonaws.com/video/dummy_video.mp4"
	file := suite.repo.FindByURI(uri)
	suite.NotNil(file)
}
