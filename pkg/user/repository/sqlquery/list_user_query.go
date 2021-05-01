package sqlquery

import (
	"fmt"
	"strings"

	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/go-pg/pg/v10/orm"
)

// ListUserQuery represents query for listing users
type ListUserQuery struct {
	scope   *orm.Query
	options domain.UserListInput
}

const (
	defaultLimit         = 50
	defaultSortColumn    = "id"
	defaultSortDirection = "desc"
)

// NewListUserQuery is a function to initialize ListUserQuery instance
func NewListUserQuery(scope *orm.Query) *ListUserQuery {
	return &ListUserQuery{scope: scope}
}

// Filter is a function to filter users by given query parameters
func (q *ListUserQuery) Filter(options domain.UserListInput) *orm.Query {
	q.options = options
	q.applySorter()
	q.filterBySearchString()
	q.filterByLimit()
	q.filterByOffset()
	return q.scope
}

// FilterBySearchString is a function to filter users by search string
func (q *ListUserQuery) filterBySearchString() {
	searchString := fmt.Sprintf("%v", q.options.SearchString)
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
	limit := q.options.Limit
	if limit == 0 {
		limit = defaultLimit
	}

	q.scope = q.scope.Limit(limit)
}

func (q *ListUserQuery) filterByOffset() {
	offset := q.options.Offset
	if q.options.Offset == 0 {
		return
	}

	q.scope = q.scope.Offset(offset)
}

func (q *ListUserQuery) applySorter() {
	var sortColumn string
	var sortDirection string

	sortColumn = q.options.SortColumn
	if sortColumn == "" {
		sortColumn = defaultSortColumn
	}

	sortDirection = q.options.SortDirection
	if sortDirection == "" {
		sortDirection = defaultSortDirection
	}

	sorter := fmt.Sprintf("%v %v NULLS LAST", sortColumn, strings.ToUpper(sortDirection))
	q.scope = q.scope.Order(sorter)
}
