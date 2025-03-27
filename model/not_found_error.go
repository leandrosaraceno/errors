package model

type NotFoundError struct {
	Msg        string
	StatusCode int
}

func NewNotFoundError(msg string) error {
	return NotFoundError{Msg: msg}
}

func (e NotFoundError) Error() string {
	return e.Msg
}
