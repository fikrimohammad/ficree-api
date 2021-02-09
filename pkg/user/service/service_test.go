package service

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/domain/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceSuite struct {
	suite.Suite
	repo *mocks.UserRepository
	svc  domain.UserService
}

func (suite *UserServiceSuite) SetupSuite() {
	suite.repo = &mocks.UserRepository{}
	suite.svc = NewUserService(suite.repo)
}

func (suite *UserServiceSuite) SetupTest() {
	suite.repo.Mock = mock.Mock{}
}

func TestUserServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for UserServiceSuite")
	}
	suite.Run(t, new(UserServiceSuite))
}

func (suite *UserServiceSuite) TestUserService_All() {
	user := &domain.User{}
	err := faker.FakeData(user)
	suite.NoError(err)

	suite.Run("when query params are valid", func() {
		queryParams := map[string]interface{}{
			"sort_column":    "name",
			"sort_direction": "desc",
			"search":         user.Title,
			"offset":         nil,
			"limit":          nil,
		}

		users := []*domain.User{user}
		suite.repo.On("List", mock.Anything).Return(users, nil)

		results, err := suite.svc.All(queryParams)
		suite.NoError(err)
		suite.NotEmpty(results)
	})

	suite.Run("when query params are invalid", func() {
		queryParams := map[string]interface{}{
			"sort_column":    10101,
			"sort_direction": 11000,
			"search":         user.Title,
			"offset":         "asasada",
			"limit":          "asasds",
		}

		results, err := suite.svc.All(queryParams)
		suite.Error(err)
		suite.Empty(results)
	})

	suite.Run("when there are error when fetching users on DB", func() {
		queryParams := map[string]interface{}{
			"search": user.Title,
		}

		suite.repo.Mock = mock.Mock{}
		suite.repo.On("List", mock.Anything).Return(nil, errors.New("dummy error"))

		results, err := suite.svc.All(queryParams)
		suite.Error(err)
		suite.Empty(results)
	})
}

func (suite *UserServiceSuite) TestUserService_Show() {
	user := &domain.User{}
	err := faker.FakeData(user)
	suite.NoError(err)

	suite.Run("when user with given id exists", func() {
		suite.repo.On("Find", user.ID).Return(user, nil)

		result, err := suite.svc.Show(user.ID)
		suite.NoError(err)
		suite.NotNil(result)
	})

	suite.Run("when there are error when find user by ID on DB", func() {
		dummyID := rand.Int()
		suite.repo.On("Find", dummyID).Return(nil, errors.New("Not Found"))

		result, err := suite.svc.Show(dummyID)
		suite.Error(err)
		suite.Nil(result)
	})
}

func (suite *UserServiceSuite) TestUserService_Create() {
	user := &domain.User{}
	_ = faker.FakeData(user)

	suite.Run("when inputs are valid", func() {
		inputs := map[string]interface{}{
			"name":         "Ficree",
			"phone_number": "62711717171",
			"email":        "ficree@gmail.com",
		}

		suite.repo.On("Create", mock.Anything).Return(user, nil)

		result, err := suite.svc.Create(inputs)
		suite.NoError(err)
		suite.NotNil(result)

		suite.repo.Mock = mock.Mock{}
	})

	suite.Run("when there are errors when storing user on DB", func() {
		inputs := map[string]interface{}{
			"name":         "Dummy",
			"phone_number": "6281111111",
			"email":        "dummy@gmail.com",
		}

		suite.repo.On("Create", mock.Anything).Return(nil, errors.New("dummy error"))

		result, err := suite.svc.Create(inputs)
		suite.Error(err)
		suite.Nil(result)

		suite.repo.Mock = mock.Mock{}
	})

	suite.Run("when inputs' type are invalid", func() {
		inputs := map[string]interface{}{
			"name":         21321231,
			"phone_number": "62711717171",
			"email":        "ficree@gmail.com",
		}

		result, err := suite.svc.Create(inputs)
		suite.Error(err)
		suite.Nil(result)
	})

	suite.Run("when inputs' values are invalid", func() {
		inputs := map[string]interface{}{
			"name":         nil,
			"phone_number": "63y6798789",
			"email":        "ficree@gmail.com",
		}

		result, err := suite.svc.Create(inputs)
		suite.Error(err)
		suite.Nil(result)
	})
}

func (suite *UserServiceSuite) TestUserService_Update() {
	user := &domain.User{}
	_ = faker.FakeData(user)
	dummyID := rand.Int()

	suite.Run("when inputs are valid", func() {
		inputs := map[string]interface{}{
			"name":         "Ficree",
			"phone_number": "62711717171",
			"email":        "ficree@gmail.com",
		}

		suite.repo.On("Update", user.ID, mock.Anything).Return(user, nil)

		result, err := suite.svc.Update(user.ID, inputs)
		suite.NoError(err)
		suite.NotNil(result)
	})

	suite.Run("when there are errors when storing user on DB", func() {
		inputs := map[string]interface{}{
			"name":         "Dummy",
			"phone_number": "6281111111",
			"email":        "dummy@gmail.com",
		}

		suite.repo.On("Update", dummyID, mock.Anything).Return(nil, errors.New("dummy error"))

		result, err := suite.svc.Update(dummyID, inputs)
		suite.Error(err)
		suite.Nil(result)
	})

	suite.Run("when inputs' type are invalid", func() {
		inputs := map[string]interface{}{
			"name":         21321231,
			"phone_number": "62711717171",
			"email":        "ficree@gmail.com",
		}

		result, err := suite.svc.Update(user.ID, inputs)
		suite.Error(err)
		suite.Nil(result)
	})

	suite.Run("when inputs' values are invalid", func() {
		inputs := map[string]interface{}{
			"email": "not_an_email",
		}

		result, err := suite.svc.Update(user.ID, inputs)
		suite.Error(err)
		suite.Nil(result)
	})
}

func (suite *UserServiceSuite) TestUserService_Destroy() {
	user := &domain.User{}
	_ = faker.FakeData(user)
	dummyID := rand.Int()

	suite.Run("when successfully delete a user", func() {
		suite.repo.On("Destroy", user.ID).Return(user, nil)

		result, err := suite.svc.Destroy(user.ID)
		suite.NoError(err)
		suite.NotNil(result)
	})

	suite.Run("when not successfully delete a user", func() {
		suite.repo.On("Destroy", dummyID).Return(nil, errors.New("dummy error"))

		result, err := suite.svc.Destroy(dummyID)
		suite.Error(err)
		suite.Nil(result)
	})
}
