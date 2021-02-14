package repository

import (
	"testing"

	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/infrastructure/database"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type UserSQLRepositorySuite struct {
	suite.Suite
	db   *pg.DB
	repo domain.UserRepository
}

func (suite *UserSQLRepositorySuite) SetupSuite() {
	suite.db = database.Load()
	suite.repo = NewSQLUserRepository(suite.db)
}

func (suite *UserSQLRepositorySuite) TearDownTest() {
	_, err := suite.db.Exec("DELETE FROM users")
	suite.NoError(err)
}

func (suite *UserSQLRepositorySuite) TearDownSuite() {
	suite.db.Close()
}

func TestUserSQLRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for UserSQLRepositorySuite")
	}
	suite.Run(t, new(UserSQLRepositorySuite))
}

func (suite *UserSQLRepositorySuite) TestUserSQLRepository_List() {
	userInput1 := &domain.User{
		Name:        "Ficree 1",
		PhoneNumber: "62811111111",
		Email:       "ficree1@gmail.com",
		Title:       "Software Engineer",
	}
	user1, err := suite.repo.Create(userInput1)
	suite.NoError(err)

	userInput2 := &domain.User{
		Name:        "Ficree 2",
		PhoneNumber: "62811111112",
		Email:       "ficree2@gmail.com",
		Title:       "Product Designer",
	}
	user2, err := suite.repo.Create(userInput2)
	suite.NoError(err)

	suite.Run("when filtered by search string", func() {
		queryParams := map[string]interface{}{
			"searchString": "Prod",
		}

		results, err := suite.repo.List(queryParams)
		suite.NoError(err)
		suite.Equal(1, len(results))
		suite.Equal(user2.GUID, results[0].GUID)
	})

	suite.Run("when filtered by limit", func() {
		queryParams := map[string]interface{}{
			"limit": 1,
		}

		results, err := suite.repo.List(queryParams)
		suite.NoError(err)
		suite.Equal(user2.GUID, results[0].GUID)
		suite.Equal(1, len(results))
	})

	suite.Run("when filtered by limit less than 1", func() {
		queryParams := map[string]interface{}{
			"limit": 0,
		}

		results, err := suite.repo.List(queryParams)
		suite.NoError(err)
		suite.Equal(user2.GUID, results[0].GUID)
		suite.Equal(user1.GUID, results[1].GUID)
		suite.Equal(2, len(results))
	})

	suite.Run("when filtered by offset", func() {
		queryParams := map[string]interface{}{
			"offset": 1,
		}

		results, err := suite.repo.List(queryParams)
		suite.NoError(err)
		suite.Equal(user1.GUID, results[0].GUID)
		suite.Equal(1, len(results))
	})

	suite.Run("when filtered by offset less than 1", func() {
		queryParams := map[string]interface{}{
			"offset": 0,
		}

		results, err := suite.repo.List(queryParams)
		suite.NoError(err)
		suite.Equal(user2.GUID, results[0].GUID)
		suite.Equal(user1.GUID, results[1].GUID)
		suite.Equal(2, len(results))
	})

	suite.Run("when order by valid column", func() {
		queryParams := map[string]interface{}{
			"sortColumn": "name",
		}

		results, err := suite.repo.List(queryParams)
		suite.NoError(err)
		suite.Equal(user2.GUID, results[0].GUID)
		suite.Equal(user1.GUID, results[1].GUID)
	})

	suite.Run("when order by valid direction", func() {
		queryParams := map[string]interface{}{
			"sortDirection": "asc",
		}

		results, err := suite.repo.List(queryParams)
		suite.NoError(err)
		suite.Equal(user1.GUID, results[0].GUID)
		suite.Equal(user2.GUID, results[1].GUID)
	})

	suite.Run("when order by invalid column", func() {
		queryParams := map[string]interface{}{
			"sortColumn": "invalid_column",
		}

		results, err := suite.repo.List(queryParams)
		suite.Error(err)
		suite.Empty(results)
	})

	suite.Run("when order by invalid direction", func() {
		queryParams := map[string]interface{}{
			"sortDirection": "invalid_direction",
		}

		results, err := suite.repo.List(queryParams)
		suite.Error(err)
		suite.Empty(results)
	})
}

func (suite *UserSQLRepositorySuite) TestUserSQLRepository_Find() {
	userInput := &domain.User{
		Name:        "Ficree",
		PhoneNumber: "62811111111",
		Email:       "ficree@gmail.com",
	}

	user, err := suite.repo.Create(userInput)
	suite.NoError(err)
	suite.NotNil(user.ID)
	suite.NotEqual(uuid.Nil, user.GUID)

	suite.Run("when user with given ID exists", func() {
		result, err := suite.repo.Find(user.ID)
		suite.NoError(err)
		suite.NotNil(result)
	})

	suite.Run("when user with given ID doesn't exist", func() {
		result, err := suite.repo.Find(-999)
		suite.Error(err)
		suite.Nil(result)
	})
}

func (suite *UserSQLRepositorySuite) TestUserSQLRepository_Create() {
	suite.Run("when inputs are valid", func() {
		params := &domain.User{
			Name:        "Ficree",
			PhoneNumber: "62811111111",
			Email:       "ficree@gmail.com",
		}

		user, err := suite.repo.Create(params)
		suite.NoError(err)
		suite.NotNil(user.ID)
		suite.NotEqual(uuid.Nil, user.GUID)
		suite.Equal(params.Name, user.Name)
		suite.Equal(params.PhoneNumber, user.PhoneNumber)
		suite.Equal(params.Email, user.Email)
	})

	suite.Run("when email input already exists", func() {
		params := &domain.User{
			Name:        "Ficree",
			PhoneNumber: "62811111112",
			Email:       "ficree@gmail.com",
		}

		user, err := suite.repo.Create(params)
		suite.Error(err)
		suite.Nil(user)
	})

	suite.Run("when phone number input already exists", func() {
		params := &domain.User{
			Name:        "Ficree",
			PhoneNumber: "62811111111",
			Email:       "ficree2@gmail.com",
		}

		user, err := suite.repo.Create(params)
		suite.Error(err)
		suite.Nil(user)
	})

	suite.Run("when name input is empty", func() {
		params := &domain.User{
			PhoneNumber: "62811111111",
			Email:       "ficree2@gmail.com",
		}

		user, err := suite.repo.Create(params)
		suite.Error(err)
		suite.Nil(user)
	})

	suite.Run("when phone number input is empty", func() {
		params := &domain.User{
			Name:  "Ficree",
			Email: "ficree2@gmail.com",
		}

		user, err := suite.repo.Create(params)
		suite.Error(err)
		suite.Nil(user)
	})

	suite.Run("when email input is empty", func() {
		params := &domain.User{
			Name:        "Ficree",
			PhoneNumber: "62811111111",
		}

		user, err := suite.repo.Create(params)
		suite.Error(err)
		suite.Nil(user)
	})
}

func (suite *UserSQLRepositorySuite) TestUserSQLRepository_Update() {
	userInput := &domain.User{
		Name:        "Ficree",
		PhoneNumber: "62811111111",
		Email:       "ficree@gmail.com",
	}

	user, err := suite.repo.Create(userInput)
	suite.NoError(err)
	suite.NotNil(user.ID)
	suite.NotEqual(uuid.Nil, user.GUID)

	userInput2 := &domain.User{
		Name:        "Ficree 2",
		PhoneNumber: "62811111112",
		Email:       "ficree2@gmail.com",
	}

	user2, err := suite.repo.Create(userInput2)
	suite.NoError(err)
	suite.NotNil(user2.ID)
	suite.NotEqual(uuid.Nil, user2.GUID)

	suite.Run("when inputs are valid", func() {
		updateInput := &domain.User{
			Name:        "Ficree Edited",
			PhoneNumber: "62811111113",
			Email:       "ficree_edited@gmail.com",
			Title:       "Software Engineer",
		}

		updatedUser, err := suite.repo.Update(user.ID, updateInput)
		suite.NoError(err)
		suite.Equal(updateInput.Name, updatedUser.Name)
		suite.Equal(updateInput.PhoneNumber, updatedUser.PhoneNumber)
		suite.Equal(updateInput.Email, updatedUser.Email)
		suite.Equal(updateInput.Title, updatedUser.Title)
	})

	suite.Run("when phone number input already used by other user", func() {
		updateInput := &domain.User{
			Name:        "Ficree Edited",
			PhoneNumber: userInput2.PhoneNumber,
			Email:       "ficree_edited@gmail.com",
		}

		updatedUser, err := suite.repo.Update(user.ID, updateInput)
		suite.Error(err)
		suite.Nil(updatedUser)
	})

	suite.Run("when email input already used by other user", func() {
		updateInput := &domain.User{
			Name:        "Ficree Edited",
			PhoneNumber: "62811111113",
			Email:       userInput2.Email,
		}

		updatedUser, err := suite.repo.Update(user.ID, updateInput)
		suite.Error(err)
		suite.Nil(updatedUser)
	})

	suite.Run("when user with given id is not found", func() {
		updateInput := &domain.User{
			Name:        "Ficree Edited",
			PhoneNumber: "62811111113",
			Email:       userInput2.Email,
		}

		updatedUser, err := suite.repo.Update(-999, updateInput)
		suite.Error(err)
		suite.Nil(updatedUser)
	})
}

func (suite *UserSQLRepositorySuite) TestUserSQLRepository_Destroy() {
	userInput := &domain.User{
		Name:        "Ficree",
		PhoneNumber: "62811111111",
		Email:       "ficree@gmail.com",
	}

	user, err := suite.repo.Create(userInput)
	suite.NoError(err)
	suite.NotNil(user.ID)
	suite.NotEqual(uuid.Nil, user.GUID)

	suite.Run("when user with given ID exists", func() {
		deletedUser, err := suite.repo.Destroy(user.ID)
		suite.NoError(err)
		suite.Nil(deletedUser)
	})

	suite.Run("when user with given ID doesn't exists", func() {
		deletedUser, err := suite.repo.Destroy(-9999)
		suite.Error(err)
		suite.Nil(deletedUser)
	})
}
