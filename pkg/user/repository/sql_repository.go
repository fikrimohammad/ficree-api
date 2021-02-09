package repository

import (
	"fmt"
	"net/http"

	"github.com/fikrimohammad/ficree-api/common/apierror"
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/pkg/user/repository/sqlquery"
	"github.com/go-pg/pg/v10"
)

// SQLUserRepository is a struct to wrap database transaction for User
type SQLUserRepository struct {
	db *pg.DB
}

// NewSQLUserRepository is a function to initialize a SQLUserRepository instance
func NewSQLUserRepository(conn *pg.DB) *SQLUserRepository {
	return &SQLUserRepository{db: conn}
}

// List is a function to fetch users
func (repo *SQLUserRepository) List(params map[string]interface{}) ([]*domain.User, error) {
	var users []*domain.User
	query := sqlquery.NewListUserQuery(repo.db.Model(&users)).Filter(params)
	err := query.Select()
	if err != nil {
		errMsg := fmt.Sprintf(domain.FindUsersError, err)
		return nil, apierror.New(http.StatusInternalServerError, errMsg)
	}

	return users, nil
}

// Find is a function to find an user by ID
func (repo *SQLUserRepository) Find(id int) (*domain.User, error) {
	var user domain.User
	query := repo.db.Model(&user).Where("id = ?", id)
	err := query.Select()
	if err != nil {
		if err == pg.ErrNoRows {
			errMsg := fmt.Sprintf(domain.FindUserByIDError, id)
			return nil, apierror.New(http.StatusNotFound, errMsg)
		}
		errMsg := fmt.Sprintf(domain.FindUserByIDError, err)
		return nil, apierror.New(http.StatusInternalServerError, errMsg)
	}

	return &user, nil
}

// Create is a function to store a new user
func (repo *SQLUserRepository) Create(newUser *domain.User) (*domain.User, error) {
	_, err := repo.db.Model(newUser).Returning("*").Insert()
	if err != nil {
		pgErr, ok := err.(pg.Error)
		if ok && pgErr.IntegrityViolation() {
			errMsg := fmt.Sprintf(domain.CreateUserError, err)
			return nil, apierror.New(http.StatusUnprocessableEntity, errMsg)
		}
		errMsg := fmt.Sprintf(domain.CreateUserError, err)
		return nil, apierror.New(http.StatusInternalServerError, errMsg)
	}

	return newUser, nil
}

// Update is a function to store a modified user
func (repo *SQLUserRepository) Update(id int, modifiedUser *domain.User) (*domain.User, error) {
	_, err := repo.db.Model(modifiedUser).
		Where("id = ?", id).
		Returning("*").
		UpdateNotZero()
	if err != nil {
		pgErr, ok := err.(pg.Error)
		if ok && pgErr.IntegrityViolation() {
			errMsg := fmt.Sprintf(domain.UpdateUserError, err)
			return nil, apierror.New(http.StatusUnprocessableEntity, errMsg)
		}
		errMsg := fmt.Sprintf(domain.UpdateUserError, err)
		return nil, apierror.New(http.StatusInternalServerError, errMsg)
	}

	return modifiedUser, nil
}

// Destroy is a function to delete an user
func (repo *SQLUserRepository) Destroy(id int) (*domain.User, error) {
	var user *domain.User
	result, err := repo.db.Model(user).
		Where("id = ?", id).
		Delete()
	if err != nil {
		errMsg := fmt.Sprintf(domain.DeleteUserError, err)
		return nil, apierror.New(http.StatusInternalServerError, errMsg)
	}

	if result.RowsAffected() <= 0 {
		errMsg := fmt.Sprintf(domain.FindUserByIDError, id)
		return nil, apierror.New(http.StatusNotFound, errMsg)
	}

	return user, nil
}
