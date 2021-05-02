package service

import (
	"testing"

	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/domain/mocks"
	"github.com/fikrimohammad/ficree-api/pkg/file/entity"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type FileServiceSuite struct {
	suite.Suite
	svc      domain.FileService
	mockRepo *mocks.FileRepository
}

func (suite *FileServiceSuite) SetupSuite() {
	suite.mockRepo = &mocks.FileRepository{}
	suite.svc = NewFileService(suite.mockRepo)
}

func TestFileServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip tests for FileServiceSuite")
	}
	suite.Run(t, new(FileServiceSuite))
}

func (suite *FileServiceSuite) TestFileServiceSuite_GenerateFileURL() {
	suite.Run("when successfully generating file URL", func() {
		params := domain.GenerateFileURLInput{
			FileFormat: "png",
			FileType:   "image",
		}

		uri := params.AsURI()
		mockFile := entity.NewAWSFile(uri)
		mockFileOutput, _ := mockFile.AsFileOutput()
		suite.mockRepo.On("FindByURI", mock.AnythingOfType("string")).Return(mockFile).Once()

		result, err := suite.svc.GetFileURL(params)
		suite.NoError(err)
		suite.Equal(mockFileOutput, result)
	})

	suite.Run("when required parameters are empty", func() {
		params := domain.GenerateFileURLInput{
			FileFormat: "",
			FileType:   "",
		}
		_, err := suite.svc.GetFileURL(params)
		suite.Error(err)
	})
}
