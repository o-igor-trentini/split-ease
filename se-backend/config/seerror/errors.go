package seerror

import "net/http"

// NewSEError cria uma instância de erro SEError.
func NewSEError(code int, message, err string, causes []Cause) SEError {
	return &seErrorImpl{
		Code:    code,
		Message: message,
		Err:     err,
		Causes:  causes,
	}
}

// NewBadRequestErr gera um erro pre-formatado para HTTP status code 400.
func NewBadRequestErr(message string) SEError {
	return NewSEError(http.StatusBadRequest, message, "bad_request", nil)
}

// NewUnauthorizedErr gera um erro pre-formatado para HTTP status code 401.
func NewUnauthorizedErr(message string) SEError {
	return NewSEError(http.StatusUnauthorized, message, "unauthorized", nil)
}

// NewForbiddenErr gera um erro pre-formatado para HTTP status code 403.
func NewForbiddenErr(message string) SEError {
	return NewSEError(http.StatusForbidden, message, "forbidden", nil)
}

// NewNotFoundErr gera um erro pre-formatado para HTTP status code 404.
func NewNotFoundErr(message string) SEError {
	return NewSEError(http.StatusNotFound, message, "not_found", nil)
}

// NewUnprocessableEntityErr gera um erro pre-formatado para HTTP status code 422.
func NewUnprocessableEntityErr(message string, causes []Cause) SEError {
	return NewSEError(http.StatusUnprocessableEntity, message, "unprocessable_entity", causes)
}

// NewsInternalServerErrorErr gera um erro pre-formatado para HTTP status code 500.
func NewsInternalServerErrorErr(message string) SEError {
	return NewSEError(http.StatusInternalServerError, message, "internal_server_error", nil)
}
