package model

import "net/http"

type BadRequestError struct {
	GenericError
}

func NewBadRequestError(msg string) error {
	return &BadRequestError{GenericError{Msg: msg, StatusCode: http.StatusBadRequest}}
}
