package apperror

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code=%d message=%s", e.Code, e.Message)
}

func NotFound(msg string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: msg}
}

func BadRequest(msg string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: msg}
}

func Unauthorized(msg string) *AppError {
	return &AppError{Code: http.StatusUnauthorized, Message: msg}
}

func Forbidden(msg string) *AppError {
	return &AppError{Code: http.StatusForbidden, Message: msg}
}

func Conflict(msg string) *AppError {
	return &AppError{Code: http.StatusConflict, Message: msg}
}

func Internal(msg string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: msg}
}
