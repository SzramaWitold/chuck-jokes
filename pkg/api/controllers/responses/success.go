package responses

type Success struct {
	Message string
}

func (r *Response) NewSuccess(message string) Success {
	return Success{Message: message}
}
