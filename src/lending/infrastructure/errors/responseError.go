package errors

type NotFoundError struct {
	message    string
	statusCode int
}

func (e NotFoundError) Error() string {
	return e.message
}

func NewNotFoundError() NotFoundError {
	return NotFoundError{message: "Not Found", statusCode: 404}
}
