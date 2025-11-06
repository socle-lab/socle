package store

import (
	"github.com/Masterminds/squirrel"
)

type Pagination struct {
	Limit  uint64
	Offset uint64
}

func ApplyPagination(q squirrel.SelectBuilder, p Pagination) squirrel.SelectBuilder {
	return q.Limit(p.Limit).Offset(p.Offset)
}
