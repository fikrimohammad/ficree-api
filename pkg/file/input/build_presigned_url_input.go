package input

import (
	"path/filepath"
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

var filePathPatternMapping = map[string]string{
	"attachment": "file_type/random/file_name.file_format",
	"image":      "file_type/random.file_format",
	"video":      "file_type/random.file_format",
}

// BuildPresignedURLInput .....
type BuildPresignedURLInput struct {
	FileName   string
	FileFormat string
	FileType   string
}

// NewBuildPresignedURLInput ......
func NewBuildPresignedURLInput(params map[string]interface{}) (*BuildPresignedURLInput, error) {
	var input BuildPresignedURLInput
	err := mapstructure.Decode(params, &input)
	return &input, err
}

// Validate ......
func (i *BuildPresignedURLInput) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(i.FileFormat, validation.Required),
		validation.Field(i.FileType, validation.Required),
	)
}

// AsURI ......
func (i *BuildPresignedURLInput) AsURI() string {
	i.preprocessParams()
	randomString := uuid.New().String()
	randomString = strings.Replace(randomString, "-", "", -1)

	uri := filePathPatternMapping[i.FileType]
	uri = strings.Replace(uri, "file_type", i.FileType, -1)
	uri = strings.Replace(uri, "file_name", i.FileName, -1)
	uri = strings.Replace(uri, "random", randomString, -1)
	uri = strings.Replace(uri, "file_format", i.FileFormat, -1)

	return uri
}

func (i *BuildPresignedURLInput) preprocessParams() {
	var fileName string
	var fileFormat string
	var fileType string

	fileType = strings.TrimSpace(i.FileType)
	fileType = strings.ToLower(fileType)
	i.FileType = fileType

	pattern := regexp.MustCompile(`/[^a-zA-Z0-9._]+/`)
	fileName = strings.ToLower(i.FileName)
	fileName = strings.TrimLeft(fileName, " ")
	fileName = strings.TrimRight(fileName, " ")
	fileName = strings.Replace(fileName, " ", "_", -1)
	fileName = pattern.ReplaceAllString(fileName, "")
	if fileName != "" {
		fileFormat = filepath.Ext(fileName)
	}
	i.FileName = fileName

	if fileFormat == "" {
		fileFormat = i.FileFormat
	}
	fileFormat = strings.TrimSpace(fileFormat)
	fileFormat = strings.ToLower(fileFormat)
	i.FileFormat = fileFormat
}
