package presenters

import "github.com/fikrimohammad/ficree-api/app/models"

type UploadPresenter struct {
	File       models.IFile
	FormatType string
}

func NewUploadPresenter(file models.IFile, formatType string) *UploadPresenter {
	presenter := UploadPresenter{File: file, FormatType: formatType}
	if presenter.FormatType == "" {
		presenter.FormatType = "format"
	}
	return &presenter
}

func (out *UploadPresenter) Result() map[string]interface{} {
	return out.format()
}

func (out *UploadPresenter) format() map[string]interface{} {
	output := map[string]interface{}{
		"upload_url": out.File.UploadURL(),
		"secure_url": out.File.PublicURL(),
	}
	return output
}
