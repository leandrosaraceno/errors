package model

type GenericError struct {
	Msg        string
	StatusCode int
}

func (e *GenericError) Error() string {
	return e.Msg
}
