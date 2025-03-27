package model

import "net/http"

type ConflictError struct {
	GenericError
}

func NewConflictError(msg string) error {
	return &ConflictError{GenericError{Msg: msg, StatusCode: http.StatusConflict}}
}
