package output

import (
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/mitchellh/mapstructure"
)

// FileOutput ......
type FileOutput struct {
	PresignedUploadURL string `mapstructure:"presigned_upload_url"`
	PublicURL          string `mapstructure:"public_url"`
}

// NewFileOutput ......
func NewFileOutput(file domain.File) (map[string]interface{}, error) {
	var result map[string]interface{}

	presignedUploadURL, err := file.UploadURL()
	if err != nil {
		return nil, err
	}

	publicURL := file.PublicURL()

	output := &FileOutput{
		PresignedUploadURL: presignedUploadURL,
		PublicURL:          publicURL,
	}

	err = mapstructure.Decode(output, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
