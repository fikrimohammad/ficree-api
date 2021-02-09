package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fikrimohammad/ficree-api/common/apiresponse"
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/gofiber/fiber/v2"
)

// FileHTTPHandler .....
type FileHTTPHandler struct {
	SVC domain.FileService
}

// HandleCreatePresignedURL ......
func (handler *FileHTTPHandler) HandleCreatePresignedURL(ctx *fiber.Ctx) error {
	var queryParams map[string]interface{}
	queryString := ctx.Context().QueryArgs().QueryString()
	err := json.Unmarshal(queryString, &queryParams)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, http.StatusUnprocessableEntity, err)
	}

	result, err := handler.SVC.BuildPresignedURL(queryParams)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, http.StatusBadRequest, err)
	}

	return apiresponse.RenderJSONSuccess(ctx, http.StatusOK, result)
}
