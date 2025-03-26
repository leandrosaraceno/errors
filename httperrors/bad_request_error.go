package httperrors

type BadRequestError struct {
	Msg string
}

func NewBadRequestError(msg string) error {
	return BadRequestError{Msg: msg}
}

func (e BadRequestError) Error() string {
	return e.Msg
}
