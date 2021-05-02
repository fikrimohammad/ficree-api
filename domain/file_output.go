package domain

// FileOutput ......
type FileOutput struct {
	PresignedUploadURL string `json:"presigned_upload_url"`
	PublicURL          string `json:"public_url"`
}
