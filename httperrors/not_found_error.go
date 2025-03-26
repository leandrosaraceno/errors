package httperrors

type NotFoundError struct {
	Msg string
}

func NewNotFoundError(msg string) error {
	return NotFoundError{Msg: msg}
}

func (e NotFoundError) Error() string {
	return e.Msg
}
