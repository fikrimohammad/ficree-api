package apiresponse

import (
	"net/http"

	"github.com/fikrimohammad/ficree-api/common/apierror"
	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error ErrResponse `json:"error,omitempty"`
}

type ErrResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Backtrace interface{} `json:"backtrace"`
}

// RenderJSONSuccess is a helper function to render JSON for successful operations
func RenderJSONSuccess(ctx *fiber.Ctx, statusCode int, data interface{}) error {
	return ctx.Status(statusCode).JSON(APIResponse{Data: &data})
}

// RenderJSONError is a helper function to render error into JSON
func RenderJSONError(ctx *fiber.Ctx, err error) error {
	var errResp ErrResponse

	statusCode := apierror.GetHTTPStatus(err)
	errResp = ErrResponse{Message: err.Error()}
	apiErr, ok := err.(*apierror.APIError)
	if !ok {
		errResp.Code = http.StatusInternalServerError
	} else {
		errResp.Code = apiErr.Code
	}

	return ctx.Status(statusCode).JSON(APIResponse{Error: errResp})
}
