package inputs

import "github.com/gofiber/fiber/v2"

// UploadPresignInput is a struct to store params for creating a new presigned URL
type UploadPresignInput struct {
	FileType   string `query:"file_type"`
	FileFormat string `query:"file_format"`
	FileName   string `query:"file_name"`
}

// NewUploadPresignInput is a function to initialize UploadPresignInput
func NewUploadPresignInput(c *fiber.Ctx) (UploadPresignInput, error) {
	input := UploadPresignInput{}
	err := c.QueryParser(&input)
	return input, err
}

// Output is a function to preprocessed query for creating presigned URL
func (i *UploadPresignInput) Output() map[string]string {
	output := make(map[string]string)
	output["file_type"] = i.FileType
	output["file_format"] = i.FileFormat
	output["file_name"] = i.FileName
	return output
}
