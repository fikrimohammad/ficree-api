package domain

// File ......
type File interface {
	PublicURL() string
	UploadURL() (string, error)
	DownloadURL() (string, error)
	AsFileOutput() (*FileOutput, error)
}

// FileRepository .....
type FileRepository interface {
	FindByURI(uri string) File
}

// FileService ......
type FileService interface {
	GetFileURL(params GenerateFileURLInput) (*FileOutput, error)
}
