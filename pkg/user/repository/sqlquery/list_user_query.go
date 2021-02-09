package sqlquery

import (
	"fmt"
	"strings"

	"github.com/go-pg/pg/v10/orm"
)

// ListUserQuery represents query for listing users
type ListUserQuery struct {
	scope   *orm.Query
	options map[string]interface{}
}

// NewListUserQuery is a function to initialize ListUserQuery instance
func NewListUserQuery(scope *orm.Query) *ListUserQuery {
	return &ListUserQuery{scope: scope}
}

// Filter is a function to filter users by given query parameters
func (q *ListUserQuery) Filter(options map[string]interface{}) *orm.Query {
	q.options = options
	q.applySorter()
	q.filterBySearchString()
	q.filterByLimit()
	q.filterByOffset()
	return q.scope
}

// FilterBySearchString is a function to filter users by search string
func (q *ListUserQuery) filterBySearchString() {
	searchString := fmt.Sprintf("%v", q.options["searchString"])
	if searchString == "" || searchString == "<nil>" {
		return
	}

	searchString = strings.ToLower(searchString)
	searchQuery := fmt.Sprintf("%%%v%%", searchString)
	q.scope = q.scope.
		WhereOr("LOWER(name) ILIKE ?", searchQuery).
		WhereOr("LOWER(title) ILIKE ?", searchQuery)
}

func (q *ListUserQuery) filterByLimit() {
	if q.options["limit"] == nil {
		q.scope = q.scope.Limit(q.defaultLimit())
		return
	}

	limit := q.options["limit"].(int)
	if limit < 1 {
		q.scope = q.scope.Limit(q.defaultLimit())
		return
	}
	q.scope = q.scope.Limit(limit)
}

func (q *ListUserQuery) filterByOffset() {
	if q.options["offset"] == nil {
		return
	}

	offset := q.options["offset"].(int)
	if offset < 1 {
		return
	}
	q.scope = q.scope.Offset(offset)
}

func (q *ListUserQuery) applySorter() {
	var sortColumn string
	var sortDirection string

	sortColumn = fmt.Sprintf("%v", q.options["sortColumn"])
	if sortColumn == "" || sortColumn == "<nil>" {
		sortColumn = "id"
	}

	sortDirection = fmt.Sprintf("%v", q.options["sortDirection"])
	if sortDirection == "" || sortDirection == "<nil>" {
		sortDirection = "desc"
	}

	sorter := fmt.Sprintf("%v %v NULLS LAST", sortColumn, strings.ToUpper(sortDirection))
	q.scope = q.scope.Order(sorter)

	if sortColumn != "id" {
		q.scope = q.scope.Order("id DESC NULLS LAST")
	}
}

func (q *ListUserQuery) defaultLimit() int {
	return 50
}
