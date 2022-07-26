package repositories

import "math"

// Pagination base return struct for Pagination
type Pagination struct {
	Page       int
	PerPage    int
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

// UpdateSettings Pagination with base setting
func (p *Pagination) UpdateSettings(page, perPage int) *Pagination {
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

// PopulateData populate pagination struct with database data
func (p *Pagination) PopulateData(totalRows int64, rows interface{}) *Pagination {
	p.TotalRows = totalRows
	p.TotalPages = int(math.Ceil(float64(totalRows) / float64(p.PerPage)))
	p.Rows = rows

	return p
}
