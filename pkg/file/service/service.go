package service

import (
	"github.com/fikrimohammad/ficree-api/domain"
)

// FileService ......
type FileService struct {
	Repo domain.FileRepository
}

// NewFileService .......
func NewFileService(repo domain.FileRepository) domain.FileService {
	return &FileService{Repo: repo}
}

// GenerateFileURL .....
func (svc *FileService) GetFileURL(params domain.GenerateFileURLInput) (*domain.FileOutput, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	file := svc.Repo.FindByURI(params.AsURI())
	result, err := file.AsFileOutput()
	if err != nil {
		return nil, err
	}

	return result, nil
}
