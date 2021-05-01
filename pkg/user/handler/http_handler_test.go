package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/fikrimohammad/ficree-api/common/apierror"
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/domain/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
)

type UserHTTPHandlerSuite struct {
	suite.Suite
	testApp     *fiber.App
	mockUserSVC *mocks.UserService
}

func (suite *UserHTTPHandlerSuite) SetupSuite() {
	mockUserSVC := &mocks.UserService{}
	userHandler := NewUserHTTPHandler(mockUserSVC)

	testApp := fiber.New()
	testApp.Get("/users", userHandler.HandleListUsers)
	testApp.Get("/users/:id", userHandler.HandleShowUser)
	testApp.Post("/users", userHandler.HandleCreateUser)
	testApp.Patch("/users/:id", userHandler.HandleUpdateUser)
	testApp.Delete("/users/:id", userHandler.HandleDestroyUser)

	suite.mockUserSVC = mockUserSVC
	suite.testApp = testApp
}

func TestUserHTTPHandlerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip tests for UserHTTPHandlerSuite")
	}
	suite.Run(t, new(UserHTTPHandlerSuite))
}

func (suite *UserHTTPHandlerSuite) TestUserHTTPHandlerSuite_HandleListUsers() {
	suite.Run("when successfully list users", func() {
		params := domain.UserListInput{
			SortColumn:    "name",
			SortDirection: "desc",
		}
		suite.mockUserSVC.On("List", params).Return([]*domain.UserCompactOutput{}, nil).Once()

		uri := fmt.Sprintf(
			"/users?sort_column=%v&sort_direction=%v",
			params.SortColumn,
			params.SortDirection,
		)
		req := httptest.NewRequest(http.MethodGet, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusOK, resp.StatusCode)
	})

	suite.Run("when the query parameters are invalid", func() {
		uri := fmt.Sprintf(
			"/users?limit=%v&offset=%v",
			"invalid%20limit",
			"invalid%20offset",
		)
		req := httptest.NewRequest(http.MethodGet, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusInternalServerError, resp.StatusCode)
	})

	suite.Run("when there is an error from service/repository layer", func() {
		dummyErr := apierror.New(http.StatusInternalServerError, "dummy error")
		params := domain.UserListInput{
			SortColumn:    faker.Word(),
			SortDirection: "asc",
		}
		suite.mockUserSVC.On("List", params).Return(nil, dummyErr).Once()

		uri := fmt.Sprintf(
			"/users?sort_column=%v&sort_direction=%v",
			params.SortColumn,
			params.SortDirection,
		)
		req := httptest.NewRequest(http.MethodGet, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusInternalServerError, resp.StatusCode)
	})
}

func (suite *UserHTTPHandlerSuite) TestUserHTTPHandlerSuite_HandleShowUser() {
	suite.Run("when successfully show an user", func() {
		dummyID := rand.Intn(10)
		suite.mockUserSVC.On("Show", dummyID).Return(&domain.UserDetailOutput{}, nil).Once()

		uri := fmt.Sprintf("/users/%v", dummyID)
		req := httptest.NewRequest(http.MethodGet, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusOK, resp.StatusCode)
	})

	suite.Run("when the path parameter is invalid", func() {
		uri := fmt.Sprintf("/users/%v", "invalid")
		req := httptest.NewRequest(http.MethodGet, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusInternalServerError, resp.StatusCode)
	})

	suite.Run("when there is an error from service/repository layer", func() {
		dummyID := rand.Intn(10)
		dummyErr := apierror.New(http.StatusNotFound, domain.FindUserByIDError)
		suite.mockUserSVC.On("Show", dummyID).Return(nil, dummyErr).Once()

		uri := fmt.Sprintf("/users/%v", dummyID)
		req := httptest.NewRequest(http.MethodGet, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusNotFound, resp.StatusCode)
	})
}

func (suite *UserHTTPHandlerSuite) TestUserHTTPHandlerSuite_HandleCreateUser() {
	suite.Run("when successfully create an user", func() {
		params := domain.UserCreateInput{
			Name:  faker.Name(),
			Email: faker.Email(),
		}
		suite.mockUserSVC.On("Create", params).Return(&domain.UserDetailOutput{}, nil).Once()

		reqBody, _ := json.Marshal(&params)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusCreated, resp.StatusCode)
	})

	suite.Run("when request body is invalid", func() {
		payload := map[string]interface{}{
			"name":  123456789,
			"email": faker.Email(),
		}
		reqBody, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusInternalServerError, resp.StatusCode)
	})

	suite.Run("when there is an error from service/repository layer", func() {
		dummyErr := apierror.New(http.StatusUnprocessableEntity, domain.CreateUserError)
		params := domain.UserCreateInput{
			Name:  faker.Name(),
			Email: faker.Email(),
		}
		suite.mockUserSVC.On("Create", params).Return(nil, dummyErr).Once()

		reqBody, _ := json.Marshal(&params)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
	})
}

func (suite *UserHTTPHandlerSuite) TestUserHTTPHandlerSuite_HandleUpdateUser() {
	suite.Run("when successfully update an user", func() {
		dummyID := rand.Int()
		params := domain.UserUpdateInput{
			Name:  faker.Name(),
			Email: faker.Email(),
		}
		suite.mockUserSVC.On("Update", dummyID, params).Return(&domain.UserDetailOutput{}, nil).Once()

		reqBody, _ := json.Marshal(&params)
		uri := fmt.Sprintf("/users/%v", dummyID)
		req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusOK, resp.StatusCode)
	})

	suite.Run("when uri path parameter is invalid", func() {
		payload := map[string]interface{}{
			"name":  123456789,
			"email": faker.Email(),
		}
		reqBody, _ := json.Marshal(payload)
		uri := fmt.Sprintf("/users/%v", "invalid")
		req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusInternalServerError, resp.StatusCode)
	})

	suite.Run("when request body is invalid", func() {
		dummyID := rand.Int()
		payload := map[string]interface{}{
			"name":  123456789,
			"email": faker.Email(),
		}
		reqBody, _ := json.Marshal(payload)
		uri := fmt.Sprintf("/users/%v", dummyID)
		req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusInternalServerError, resp.StatusCode)
	})

	suite.Run("when there is an error from service/repository layer", func() {
		dummyID := rand.Int()
		dummyErr := apierror.New(http.StatusUnprocessableEntity, domain.UpdateUserError)
		params := domain.UserUpdateInput{
			Name:  faker.Name(),
			Email: faker.Email(),
		}
		suite.mockUserSVC.On("Update", dummyID, params).Return(nil, dummyErr).Once()

		reqBody, _ := json.Marshal(&params)
		uri := fmt.Sprintf("/users/%v", dummyID)
		req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
	})
}

func (suite *UserHTTPHandlerSuite) TestUserHTTPHandlerSuite_HandleDestroyUser() {
	suite.Run("when successfully destroy an user", func() {
		dummyID := rand.Int()
		suite.mockUserSVC.On("Destroy", dummyID).Return(&domain.UserDetailOutput{}, nil).Once()

		uri := fmt.Sprintf("/users/%v", dummyID)
		req := httptest.NewRequest(http.MethodDelete, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusOK, resp.StatusCode)
	})

	suite.Run("when uri path parameter is invalid", func() {
		uri := fmt.Sprintf("/users/%v", "invalid")
		req := httptest.NewRequest(http.MethodDelete, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusInternalServerError, resp.StatusCode)
	})

	suite.Run("when there is an error from service/repository layer", func() {
		dummyID := rand.Int()
		dummyErr := apierror.New(http.StatusNotFound, domain.FindUserByIDError)
		suite.mockUserSVC.On("Destroy", dummyID).Return(nil, dummyErr).Once()

		uri := fmt.Sprintf("/users/%v", dummyID)
		req := httptest.NewRequest(http.MethodDelete, uri, nil)

		resp, err := suite.testApp.Test(req)
		suite.NoError(err)
		suite.Equal(http.StatusNotFound, resp.StatusCode)
	})
}
