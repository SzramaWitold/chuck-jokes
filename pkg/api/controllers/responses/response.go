package responses

type Response struct{}

func NewResponse() *Response {
	return &Response{}
}

type Pagination[T interface{}] struct {
	Page       int
	PerPage    int
	TotalRows  int64
	TotalPages int
	Rows       []T
}

func ResponsePagination[T interface{}](page int, perPage int, totalRows int64, totalPages int, rows []T) *Pagination[T] {
	return &Pagination[T]{Page: page, PerPage: perPage, TotalRows: totalRows, TotalPages: totalPages, Rows: rows}
}
