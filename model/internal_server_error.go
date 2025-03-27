package model

import "net/http"

type InternalServerError struct {
	GenericError
}

func NewInternalServerError(msg string) error {
	return &InternalServerError{GenericError{Msg: msg, StatusCode: http.StatusInternalServerError}}
}
