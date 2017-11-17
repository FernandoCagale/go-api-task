package models

import (
	"strconv"

	"github.com/labstack/echo"
)

type Query struct {
	Search string
	Limit  int
	Offset int
}

func NewQuery(c echo.Context) *Query {
	search := c.QueryParam("search")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	return &Query{search, limit, offset}
}

func (q *Query) IsEmpty() bool {
	return q.Search == ""
}

func (q *Query) GetLimit() int {
	if q.Limit == 0 {
		return 10
	}
	return q.Limit
}
