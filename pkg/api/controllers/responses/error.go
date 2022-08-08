package responses

type Error struct {
	Message string
}

func (r *Response) NewError(err error) Error {
	return Error{
		Message: err.Error(),
	}
}

func (r *Response) NewErrorsCollection(errors []error) []Error {
	var collection []Error

	for _, err := range errors {
		collection = append(collection, r.NewError(err))
	}

	return collection
}
