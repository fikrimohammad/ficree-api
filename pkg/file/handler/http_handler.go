package handler

import (
	"encoding/json"
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

// HandleCreatePresignedURL ...
func (handler *FileHTTPHandler) HandleCreatePresignedURL(ctx *fiber.Ctx) error {
	var queryParams map[string]interface{}
	queryString := ctx.Context().QueryArgs().QueryString()
	err := json.Unmarshal(queryString, &queryParams)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	result, err := handler.SVC.BuildPresignedURL(queryParams)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	return apiresponse.RenderJSONSuccess(ctx, http.StatusOK, result)
}
