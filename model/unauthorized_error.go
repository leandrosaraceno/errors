package model

import "net/http"

type UnauthorizedError struct {
	genericError
}

func NewUnauthorizedError(msg string) error {
	return UnauthorizedError{genericError{Msg: msg, StatusCode: http.StatusUnauthorized}}
}
