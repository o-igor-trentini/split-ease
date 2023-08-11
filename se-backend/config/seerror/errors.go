package seerror

import (
	"net/http"
)

const (
	ErrBadRequest          = "bad_request"
	ErrUnauthorized        = "unauthorized"
	ErrForbidden           = "forbidden"
	ErrNotFound            = "not_found"
	ErrUnprocessableEntity = "unprocessable_entity"
	ErrInternalServerError = "internal_server_error"
)

// NewBadRequestErr gera um erro pre-formatado para HTTP status code 400.
func NewBadRequestErr(message string, source error) SEError {
	return New(http.StatusBadRequest, message, ErrBadRequest, nil, source)
}

// NewUnauthorizedErr gera um erro pre-formatado para HTTP status code 401.
func NewUnauthorizedErr(message string, source error) SEError {
	return New(http.StatusUnauthorized, message, ErrUnauthorized, nil, source)
}

// NewForbiddenErr gera um erro pre-formatado para HTTP status code 403.
func NewForbiddenErr(message string, source error) SEError {
	return New(http.StatusForbidden, message, ErrForbidden, nil, source)
}

// NewNotFoundErr gera um erro pre-formatado para HTTP status code 404.
func NewNotFoundErr(message string, source error) SEError {
	return New(http.StatusNotFound, message, ErrNotFound, nil, source)
}

// NewUnprocessableEntityErr gera um erro pre-formatado para HTTP status code 422.
func NewUnprocessableEntityErr(message string, source error, causes []Cause) SEError {
	return New(http.StatusUnprocessableEntity, message, ErrUnprocessableEntity, causes, source)
}

// NewsInternalServerErrorErr gera um erro pre-formatado para HTTP status code 500.
func NewsInternalServerErrorErr(message string, source error) SEError {
	return New(http.StatusInternalServerError, message, ErrInternalServerError, nil, source)
}
