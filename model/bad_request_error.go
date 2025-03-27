package model

import "net/http"

type BadRequestError struct {
	genericError
}

func NewBadRequestError(msg string) error {
	return BadRequestError{genericError{Msg: msg, StatusCode: http.StatusBadRequest}}
}
