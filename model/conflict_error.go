package model

import "net/http"

type ConflictError struct {
	genericError
}

func NewConflictError(msg string) error {
	return ConflictError{genericError{Msg: msg, StatusCode: http.StatusConflict}}
}
