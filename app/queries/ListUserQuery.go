package queries

import (
	"fmt"

	"github.com/fikrimohammad/ficree-api/app/models"
	"gorm.io/gorm"
)

type ListUserQuery struct {
	conn *gorm.DB
}

func NewListUserQuery(conn *gorm.DB) *ListUserQuery {
	return &ListUserQuery{conn: conn}
}

func (q *ListUserQuery) Filter(options map[string]interface{}) *gorm.DB {
	scope := q.defaultScope()
	if len(options) != 0 {
		scope = q.filterByName(scope, options["name"])
		scope = q.filterByLimit(scope, options["limit"])
	}
	return scope
}

func (q *ListUserQuery) filterByName(scope *gorm.DB, name interface{}) *gorm.DB {
	if name == nil || name == "" {
		return scope
	}
	value := fmt.Sprintf("%%%v%%", name)
	return scope.Where("name ILIKE ?", value)
}

func (q *ListUserQuery) filterByLimit(scope *gorm.DB, limit interface{}) *gorm.DB {
	if limit == nil || limit == "" {
		return scope.Limit(q.defaultLimit())
	}
	return scope.Limit(limit.(int))
}

func (q *ListUserQuery) defaultScope() *gorm.DB {
	return q.conn.Model(&models.User{})
}

func (q *ListUserQuery) defaultLimit() int {
	return 50
}
