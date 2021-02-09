package service

import (
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/pkg/file/input"
	"github.com/fikrimohammad/ficree-api/pkg/file/output"
)

// FileService ......
type FileService struct {
	Repo domain.FileRepository
}

// NewFileService .......
func NewFileService(repo domain.FileRepository) domain.FileService {
	return &FileService{Repo: repo}
}

// BuildPresignedURL .....
func (svc *FileService) BuildPresignedURL(params map[string]interface{}) (map[string]interface{}, error) {
	input, err := input.NewBuildPresignedURLInput(params)
	if err != nil {
		return nil, err
	}

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	file, err := svc.Repo.FindByURI(input.AsURI())
	if err != nil {
		return nil, err
	}

	result, err := output.NewFileOutput(file)
	if err != nil {
		return nil, err
	}

	return result, nil
}
