package httperrors

type UnauthorizedError struct {
	Msg string
}

func NewUnauthorizedError(msg string) error {
	return UnauthorizedError{Msg: msg}
}

func (e UnauthorizedError) Error() string {
	return e.Msg
}
