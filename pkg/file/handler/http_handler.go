package handler

import (
	"net/http"

	"github.com/fikrimohammad/ficree-api/common/apiresponse"
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/gofiber/fiber/v2"
)

// FileHTTPHandler ...
type FileHTTPHandler struct {
	SVC domain.FileService
}

// NewFileHTTPHandler ...
func NewFileHTTPHandler(svc domain.FileService) FileHTTPHandler {
	return FileHTTPHandler{SVC: svc}
}

// HandleGenerateFileURL ...
func (handler *FileHTTPHandler) HandleGetFileURL(ctx *fiber.Ctx) error {
	var queryParams domain.GenerateFileURLInput
	err := ctx.QueryParser(&queryParams)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	result, err := handler.SVC.GetFileURL(queryParams)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	return apiresponse.RenderJSONSuccess(ctx, http.StatusOK, result)
}
