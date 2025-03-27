package model

type UnauthorizedError struct {
	Msg        string
	StatusCode int
}

func NewUnauthorizedError(msg string) error {
	return UnauthorizedError{Msg: msg}
}

func (e UnauthorizedError) Error() string {
	return e.Msg
}
