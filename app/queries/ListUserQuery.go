package queries

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fikrimohammad/ficree-api/app/models"
	"gorm.io/gorm"
)

type ListUserQuery struct {
	scope *gorm.DB
}

func NewListUserQuery(scope *gorm.DB) *ListUserQuery {
	return &ListUserQuery{scope: scope}
}

func (q *ListUserQuery) Filter(options map[string]interface{}) *gorm.DB {
	q.setDefaultScope()
	q.applyFilters(options)
	return q.scope
}

func (q *ListUserQuery) applyFilters(options map[string]interface{}) {
	for key, val := range options {
		if val == nil || val == "" {
			continue
		}

		methodName := fmt.Sprintf("FilterBy%v", strings.Title(key))
		methodParams := []reflect.Value{}
		methodParams = append(methodParams, reflect.ValueOf(val))
		method := reflect.ValueOf(q).MethodByName(methodName)
		method.Call(methodParams)
	}
}

func (q *ListUserQuery) FilterByName(name string) {
	value := fmt.Sprintf("%%%v%%", name)
	q.scope = q.scope.Where("name ILIKE ?", value)
}

func (q *ListUserQuery) FilterByLimit(limit int) {
	q.scope = q.scope.Limit(limit)
}

func (q *ListUserQuery) setDefaultScope() {
	q.scope = q.scope.Model(&models.User{}).Limit(q.defaultLimit())
}

func (q *ListUserQuery) defaultLimit() int {
	return 50
}
