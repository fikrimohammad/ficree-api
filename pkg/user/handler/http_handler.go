package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fikrimohammad/ficree-api/common/apiresponse"
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/gofiber/fiber/v2"
)

// UserHTTPHandler represents HTTP handler for users
type UserHTTPHandler struct {
	SVC domain.UserService
}

// NewUserHTTPHandler is a function to initialize UserHTTPHandler instance
func NewUserHTTPHandler(svc domain.UserService) UserHTTPHandler {
	return UserHTTPHandler{SVC: svc}
}

// HandleListUsers is a handler to list users by supported query parameters
func (handler *UserHTTPHandler) HandleListUsers(ctx *fiber.Ctx) error {
	var queryParams map[string]interface{}
	queryString := ctx.Context().QueryArgs().QueryString()
	err := json.Unmarshal(queryString, &queryParams)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	results, err := handler.SVC.All(queryParams)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	return apiresponse.RenderJSONSuccess(ctx, http.StatusOK, results)
}

// HandleShowUser is an API to show an user
func (handler *UserHTTPHandler) HandleShowUser(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	result, err := handler.SVC.Show(userID)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	return apiresponse.RenderJSONSuccess(ctx, http.StatusOK, result)
}

// HandleCreateUser is an API to create an user
func (handler *UserHTTPHandler) HandleCreateUser(ctx *fiber.Ctx) error {
	var params map[string]interface{}
	err := json.Unmarshal(ctx.Body(), &params)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	user, err := handler.SVC.Create(params)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	return apiresponse.RenderJSONSuccess(ctx, http.StatusCreated, user)
}

// HandleUpdateUser is an API to update an user
func (handler *UserHTTPHandler) HandleUpdateUser(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	var params map[string]interface{}
	err = json.Unmarshal(ctx.Body(), &params)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	user, err := handler.SVC.Update(userID, params)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	return apiresponse.RenderJSONSuccess(ctx, http.StatusOK, user)
}

// HandleDestroyUser is an API to delete an user
func (handler *UserHTTPHandler) HandleDestroyUser(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	user, err := handler.SVC.Destroy(userID)
	if err != nil {
		return apiresponse.RenderJSONError(ctx, err)
	}

	return apiresponse.RenderJSONSuccess(ctx, http.StatusOK, user)
}
