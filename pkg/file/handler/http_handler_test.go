package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/domain/mocks"
	"github.com/fikrimohammad/ficree-api/pkg/file/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
)

type FileHTTPHandlerSuite struct {
	suite.Suite
	testApp     *fiber.App
	mockFileSVC *mocks.FileService
}

func (suite *FileHTTPHandlerSuite) SetupSuite() {
	mockFileSVC := &mocks.FileService{}
	fileHandler := NewFileHTTPHandler(mockFileSVC)

	testApp := fiber.New()
	testApp.Get("/files/presign", fileHandler.HandleGetFileURL)

	suite.mockFileSVC = mockFileSVC
	suite.testApp = testApp
}

func TestFileHTTPHandlerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip tests for FileHTTPHandlerSuite")
	}
	suite.Run(t, new(FileHTTPHandlerSuite))
}

func (suite *FileHTTPHandlerSuite) TestFileHTTPHandlerSuite_HandleGetFileURL() {
	suite.Run("when successfully get file URL", func() {
		params := domain.GenerateFileURLInput{
			FileFormat: "png",
			FileType:   "image",
		}

		mockFile := entity.NewAWSFile(params.AsURI())
		mockFileOutput, _ := mockFile.AsFileOutput()
		suite.mockFileSVC.On("GetFileURL", params).Return(mockFileOutput, nil).Once()

		uri := fmt.Sprintf(
			"/files/presign?file_format=%v&file_type=%v",
			params.FileFormat,
			params.FileType,
		)
		req := httptest.NewRequest(http.MethodGet, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusOK, resp.StatusCode)
	})

	suite.Run("when there is an error from service/repository layer", func() {
		params := domain.GenerateFileURLInput{
			FileFormat: "mp4",
			FileType:   "video",
		}
		dummyErr := errors.New("dummy error")
		suite.mockFileSVC.On("GetFileURL", params).Return(nil, dummyErr).Once()

		uri := fmt.Sprintf(
			"/files/presign?file_format=%v&file_type=%v",
			params.FileFormat,
			params.FileType,
		)
		req := httptest.NewRequest(http.MethodGet, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusInternalServerError, resp.StatusCode)
	})
}
