package domain

// File ......
type File interface {
	PublicURL() string
	UploadURL() (string, error)
	DownloadURL() (string, error)
}

// FileRepository .....
type FileRepository interface {
	FindByURI(uri string) (File, error)
}

// FileService ......
type FileService interface {
	BuildPresignedURL(params map[string]interface{}) (map[string]interface{}, error)
}
