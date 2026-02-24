package payspace

import (
	"fmt"
	"net/url"
	"strings"
)

// QueryParams builds OData query parameters for API requests.
type QueryParams struct {
	filter  string
	top     *int
	skip    *int
	orderBy string
	selekt  []string // "select" is a Go keyword
	expand  []string
	count   bool
	extra   map[string]string
}

// NewQuery creates a new OData query builder.
func NewQuery() *QueryParams {
	return &QueryParams{}
}

// Filter sets the $filter expression.
//
//	q.Filter("EmployeeNumber eq 'EMP001'")
func (q *QueryParams) Filter(expr string) *QueryParams {
	q.filter = expr
	return q
}

// Top sets the maximum number of records to return (max 100).
func (q *QueryParams) Top(n int) *QueryParams {
	if n > maxPageSize {
		n = maxPageSize
	}
	q.top = &n
	return q
}

// Skip sets the number of records to skip for pagination.
func (q *QueryParams) Skip(n int) *QueryParams {
	q.skip = &n
	return q
}

// OrderBy sets the sort expression.
//
//	q.OrderBy("Surname asc")
func (q *QueryParams) OrderBy(expr string) *QueryParams {
	q.orderBy = expr
	return q
}

// Select specifies which fields to return.
//
//	q.Select("EmployeeNumber", "FirstName", "Surname")
func (q *QueryParams) Select(fields ...string) *QueryParams {
	q.selekt = fields
	return q
}

// Expand includes related entities in the response.
//
//	q.Expand("OrganizationGroups")
func (q *QueryParams) Expand(fields ...string) *QueryParams {
	q.expand = fields
	return q
}

// Count requests the total count in the response (@odata.count).
func (q *QueryParams) Count() *QueryParams {
	q.count = true
	return q
}

// Param sets an arbitrary query parameter (e.g., frequency, period).
func (q *QueryParams) Param(key, value string) *QueryParams {
	if q.extra == nil {
		q.extra = make(map[string]string)
	}
	q.extra[key] = value
	return q
}

// Encode serializes the query parameters into a URL query string.
func (q *QueryParams) Encode() string {
	if q == nil {
		return ""
	}
	v := url.Values{}
	if q.filter != "" {
		v.Set("$filter", q.filter)
	}
	if q.top != nil {
		v.Set("$top", fmt.Sprintf("%d", *q.top))
	}
	if q.skip != nil {
		v.Set("$skip", fmt.Sprintf("%d", *q.skip))
	}
	if q.orderBy != "" {
		v.Set("$orderby", q.orderBy)
	}
	if len(q.selekt) > 0 {
		v.Set("$select", strings.Join(q.selekt, ","))
	}
	if len(q.expand) > 0 {
		v.Set("$expand", strings.Join(q.expand, ","))
	}
	if q.count {
		v.Set("$count", "true")
	}
	for key, val := range q.extra {
		v.Set(key, val)
	}
	return v.Encode()
}
