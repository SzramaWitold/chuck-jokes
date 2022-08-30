package responses

type Error struct {
	Message string
}

func (r *DefaultResponseHandler) NewError(err error) Error {
	return Error{
		Message: err.Error(),
	}
}

func (r *DefaultResponseHandler) NewErrorsCollection(errors []error) []Error {
	var collection []Error

	for _, err := range errors {
		collection = append(collection, r.NewError(err))
	}

	return collection
}
