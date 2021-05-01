package service

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/fikrimohammad/ficree-api/common/apierror"
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/domain/mocks"
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

	suite.Run("when successfully fetch all users", func() {
		queryParams := domain.UserListInput{
			SortColumn:    "name",
			SortDirection: "desc",
			SearchString:  user.Title,
			Offset:        0,
			Limit:         0,
		}

		users := []*domain.User{user}
		suite.repo.On("List", queryParams).Return(users, nil).Once()

		results, err := suite.svc.List(queryParams)
		suite.NoError(err)
		suite.IsType([]*domain.UserCompactOutput{}, results)
		suite.Equal(len(users), len(results))
	})

	suite.Run("when there are error when fetching users on DB", func() {
		queryParams := domain.UserListInput{
			SortColumn:    "invalid column",
			SortDirection: "invalid direction",
			SearchString:  user.Title,
			Offset:        0,
			Limit:         0,
		}

		dummyErr := errors.New("dummy error")
		suite.repo.On("List", queryParams).Return(nil, dummyErr).Once()

		results, err := suite.svc.List(queryParams)
		suite.EqualError(err, dummyErr.Error())
		suite.Empty(results)
	})
}

func (suite *UserServiceSuite) TestUserService_Show() {
	user := &domain.User{}
	err := faker.FakeData(user)
	suite.NoError(err)

	suite.Run("when user with given id exists", func() {
		suite.repo.On("Find", user.ID).Return(user, nil).Once()

		result, err := suite.svc.Show(user.ID)
		suite.NoError(err)
		suite.NotNil(result)
	})

	suite.Run("when there are error when find user by ID on DB", func() {
		dummyID := rand.Int()
		dummyErr := apierror.New(http.StatusNotFound, fmt.Sprintf(domain.FindUserByIDError, dummyID))
		suite.repo.On("Find", dummyID).Return(nil, dummyErr).Once()

		result, err := suite.svc.Show(dummyID)
		suite.EqualError(err, dummyErr.Error())
		suite.Nil(result)
	})
}

func (suite *UserServiceSuite) TestUserService_Create() {
	user := &domain.User{}
	_ = faker.FakeData(user)

	suite.Run("when successfully create a user", func() {
		input := domain.UserCreateInput{
			Name:        "Ficree",
			PhoneNumber: "6285894298777",
			Email:       faker.Email(),
		}

		suite.repo.On("Create", input.AsUser()).Return(user, nil).Once()

		result, err := suite.svc.Create(input)
		suite.NoError(err)
		suite.NotNil(result)
	})

	suite.Run("when there are errors when storing user on DB", func() {
		input := domain.UserCreateInput{
			Name:        "Ficree",
			PhoneNumber: "6285894298777",
			Email:       faker.Email(),
		}

		dummyErr := apierror.New(http.StatusUnprocessableEntity, domain.CreateUserError)
		suite.repo.On("Create", input.AsUser()).Return(nil, dummyErr).Once()

		result, err := suite.svc.Create(input)
		suite.EqualError(err, dummyErr.Error())
		suite.Nil(result)
	})

	suite.Run("when input is invalid", func() {
		inputs := domain.UserCreateInput{
			Name:        "",
			PhoneNumber: "63y6798789",
			Email:       "ficree@gmail.com",
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

	suite.Run("when successfully update a user", func() {
		input := domain.UserUpdateInput{
			Name:        "Ficree",
			PhoneNumber: "628581234567",
			Email:       faker.Email(),
		}

		suite.repo.On("Update", user.ID, input.AsUser()).Return(user, nil).Once()

		result, err := suite.svc.Update(user.ID, input)
		suite.NoError(err)
		suite.NotNil(result)
	})

	suite.Run("when there are errors when storing user on DB", func() {
		input := domain.UserUpdateInput{
			Name:        "Ficree",
			PhoneNumber: "628581234567",
			Email:       faker.Email(),
		}

		dummyErr := apierror.New(http.StatusUnprocessableEntity, domain.UpdateUserError)
		suite.repo.On("Update", dummyID, input.AsUser()).Return(nil, dummyErr).Once()

		result, err := suite.svc.Update(dummyID, input)
		suite.EqualError(err, dummyErr.Error())
		suite.Nil(result)
	})

	suite.Run("when input is invalid", func() {
		input := domain.UserUpdateInput{
			PhoneNumber: "+68986899",
		}

		result, err := suite.svc.Update(user.ID, input)
		suite.Error(err)
		suite.Nil(result)
	})
}

func (suite *UserServiceSuite) TestUserService_Destroy() {
	user := &domain.User{}
	_ = faker.FakeData(user)
	dummyID := rand.Int()

	suite.Run("when successfully delete a user", func() {
		suite.repo.On("Destroy", user.ID).Return(user, nil).Once()

		result, err := suite.svc.Destroy(user.ID)
		suite.NoError(err)
		suite.NotNil(result)
	})

	suite.Run("when not successfully delete a user", func() {
		suite.repo.On("Destroy", dummyID).Return(nil, errors.New("dummy error")).Once()

		result, err := suite.svc.Destroy(dummyID)
		suite.Error(err)
		suite.Nil(result)
	})
}
