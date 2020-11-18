package services

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/google/uuid"
)

var fileManagers = map[string]map[string]interface{}{
	"attachment": {
		"path":      "file_type/random/file_name.file_format",
		"file_type": "attachment",
	},
	"image": {
		"path":      "file_type/random.file_format",
		"file_type": "image",
	},
	"video": {
		"path":      "file_type/random.file_format",
		"file_type": "video",
	},
}

type UploadService struct{}

func NewUploadService() *UploadService {
	return &UploadService{}
}

func (svc *UploadService) BuildPresignedURL(params map[string]string) (models.IFile, error) {
	uploadParams := buildParams(params)
	filePath := buildFilePath(uploadParams)
	file, err := models.NewFile(filePath)
	return file, err
}

func buildParams(params map[string]string) map[string]string {
	var fileName string
	var fileFormat string

	params["file_type"] = strings.TrimSpace(params["file_type"])
	params["file_type"] = strings.ToLower(params["file_type"])

	pattern := regexp.MustCompile(`/[^a-zA-Z0-9._]+/`)
	fileName = params["file_name"]
	fileName = strings.ToLower(fileName)
	fileName = strings.TrimLeft(fileName, " ")
	fileName = strings.TrimRight(fileName, " ")
	fileName = strings.Replace(fileName, " ", "_", -1)
	fileName = pattern.ReplaceAllString(fileName, "")
	if fileName != "" {
		fileFormat = filepath.Ext(fileName)
	}
	params["file_name"] = fileName

	if fileFormat == "" {
		fileFormat = params["file_format"]
	}
	fileFormat = strings.TrimSpace(fileFormat)
	fileFormat = strings.ToLower(fileFormat)
	params["file_format"] = fileFormat

	return params
}

func buildFilePath(params map[string]string) string {
	randomUUID, _ := uuid.NewRandom()
	randomString := randomUUID.String()
	randomString = strings.Replace(randomString, "-", "", -1)

	fileManager := fileManagers[params["file_type"]]
	path := fmt.Sprintf("%v", fileManager["path"])
	path = strings.Replace(path, "file_type", params["file_type"], -1)
	path = strings.Replace(path, "file_name", params["file_name"], -1)
	path = strings.Replace(path, "random", randomString, -1)
	path = strings.Replace(path, "file_format", params["file_format"], -1)

	return path
}
