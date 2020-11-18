package interfaces

import "github.com/fikrimohammad/ficree-api/app/models"

type IUploadService interface {
	BuildPresignedURL(params map[string]string) (models.IFile, error)
}
