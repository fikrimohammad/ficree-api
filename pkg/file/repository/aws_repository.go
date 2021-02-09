package repository

import (
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/pkg/file/entity"
)

// AWSFileRepository ......
type AWSFileRepository struct{}

// NewAWSFileRepository .....
func NewAWSFileRepository() domain.FileRepository {
	return &AWSFileRepository{}
}

// FindByURI .....
func (repo *AWSFileRepository) FindByURI(uri string) (domain.File, error) {
	file, err := entity.NewAWSFile(uri)
	if err != nil {
		return nil, err
	}

	return file, nil
}
