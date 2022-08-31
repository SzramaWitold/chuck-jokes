package gorm

import "math"

// Pagination base return struct for Pagination
type Pagination[T interface{}] struct {
	Page       int
	PerPage    int
	TotalRows  int64
	TotalPages int
	Rows       []T
}

func NewPagination[T interface{}]() *Pagination[T] {
	return &Pagination[T]{}
}

// UpdateSettings Pagination with base setting
func (p *Pagination[T]) UpdateSettings(page, perPage int) *Pagination[T] {
	if page == 0 {
		p.Page = 1
	} else {
		p.Page = page
	}

	switch {
	case perPage > 100:
		p.PerPage = 100
	case perPage <= 0:
		p.PerPage = 10
	default:
		p.PerPage = perPage
	}

	return p
}

// PopulateData populate Pagination struct with database data
func (p *Pagination[T]) PopulateData(totalRows int64, rows []T) *Pagination[T] {
	p.TotalRows = totalRows
	p.TotalPages = int(math.Ceil(float64(totalRows) / float64(p.PerPage)))
	p.Rows = rows

	return p
}
