package model

import "net/http"

type NotFoundError struct {
	GenericError
}

func NewNotFoundError(msg string) error {
	return &NotFoundError{GenericError{Msg: msg, StatusCode: http.StatusNotFound}}
}
