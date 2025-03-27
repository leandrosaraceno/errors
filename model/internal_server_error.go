package model

import "net/http"

type InternalServerError struct {
	genericError
}

func NewInternalServerError(msg string) error {
	return InternalServerError{genericError{Msg: msg, StatusCode: http.StatusInternalServerError}}
}
