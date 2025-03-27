package model

import "net/http"

type UnauthorizedError struct {
	GenericError
}

func NewUnauthorizedError(msg string) error {
	return &UnauthorizedError{GenericError{Msg: msg, StatusCode: http.StatusUnauthorized}}
}
