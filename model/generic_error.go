package model

type GenericError interface {
	GetMessage() string
	GetStatusCode() int
}

type genericError struct {
	Msg        string
	StatusCode int
}

func (e genericError) Error() string {
	return e.Msg
}

func (e genericError) GetMessage() string {
	return e.Msg
}

func (e genericError) GetStatusCode() int {
	return e.StatusCode
}
