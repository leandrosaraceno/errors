package httperrors

type InternalServerError struct {
	Msg string
}

func NewInternalServerError(msg string) error {
	return InternalServerError{Msg: msg}
}

func (e InternalServerError) Error() string {
	return e.Msg
}
