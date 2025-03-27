package model

import "net/http"

type NotFoundError struct {
	genericError
}

func NewNotFoundError(msg string) error {
	return NotFoundError{genericError{Msg: msg, StatusCode: http.StatusNotFound}}
}
