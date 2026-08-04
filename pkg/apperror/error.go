package apperror

import (
	"evara-backend/pkg/response"
	"net/http"
)

type Error struct {
	HTTPStatus int
	Code       string
	Message    string
}

func (e *Error) Error() string {
	return e.Message
}

func BadRequest(message string) error {
	return &Error{
		Code:       response.CodeBadRequest,
		Message:    message,
		HTTPStatus: http.StatusBadRequest,
	}
}

func Validation(message string) error {
	return &Error{
		Code:       response.CodeValidation,
		Message:    message,
		HTTPStatus: http.StatusBadRequest,
	}
}

func Unauthorized(message string) error {
	return &Error{
		Code:       response.CodeUnauthorized,
		Message:    message,
		HTTPStatus: http.StatusUnauthorized,
	}
}

func Forbidden(message string) error {
	return &Error{
		Code:       response.CodeForbidden,
		Message:    message,
		HTTPStatus: http.StatusForbidden,
	}
}

func NotFound(message string) error {
	return &Error{
		Code:       response.CodeNotFound,
		Message:    message,
		HTTPStatus: http.StatusNotFound,
	}
}

func Conflict(message string) error {
	return &Error{
		Code:       response.CodeConflict,
		Message:    message,
		HTTPStatus: http.StatusConflict,
	}
}

func Internal(message string) error {
	return &Error{
		Code:       response.CodeInternal,
		Message:    "Internal server error",
		HTTPStatus: http.StatusInternalServerError,
	}
}