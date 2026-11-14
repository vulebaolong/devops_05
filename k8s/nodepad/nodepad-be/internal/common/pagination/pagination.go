package pagination

import (
	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
)

const (
	DefaultPage     = 1
	DefaultPageSize = 10
)

func Normalize(page, pageSize int) (int, int, int) {
	if page <= 0 {
		page = DefaultPage
	}

	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}

	offset := (page - 1) * pageSize

	return int(page), int(pageSize), int(offset)
}

// ListOption configures list queries.
type ListOption func(*ListOptions)

// ListOptions are list query options.
type ListOptions struct {
	Filter  filtering.Filter
	OrderBy ordering.OrderBy
	Offset  int
	Limit   int
}

// ListFilter sets a standard AIP filter.
func ListFilter(filter filtering.Filter) ListOption {
	return func(o *ListOptions) {
		o.Filter = filter
	}
}

// ListOrderBy sets a standard AIP order_by value.
func ListOrderBy(orderBy ordering.OrderBy) ListOption {
	return func(o *ListOptions) {
		o.OrderBy = orderBy
	}
}

// ListOffset sets an offset.
func ListOffset(offset int) ListOption {
	return func(o *ListOptions) {
		o.Offset = offset
	}
}

// ListLimit sets a limit.
func ListLimit(limit int) ListOption {
	return func(o *ListOptions) {
		o.Limit = limit
	}
}
