package responses

type Success struct {
	Message string
}

func (r *DefaultResponseHandler) NewSuccess(message string) Success {
	return Success{Message: message}
}
