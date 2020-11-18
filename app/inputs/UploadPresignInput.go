package inputs

import "github.com/gofiber/fiber/v2"

type UploadPresignInput struct {
	FileType   string `query:"file_type"`
	FileFormat string `query:"file_format"`
	FileName   string `query:"file_name"`
}

func NewUploadPresignInput(c *fiber.Ctx) (UploadPresignInput, error) {
	input := UploadPresignInput{}
	err := c.QueryParser(&input)
	return input, err
}

func (i *UploadPresignInput) Output() map[string]string {
	output := make(map[string]string)
	output["file_type"] = i.FileType
	output["file_format"] = i.FileFormat
	output["file_name"] = i.FileName
	return output
}
