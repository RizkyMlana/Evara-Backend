package apperror

import "net/http"

type Error struct {
	Code       string
	Message    string
	HTTPStatus int
}

func (e *Error) Error() string {
	return e.Message
}

func BadRequest(message string) error {
	return &Error{
		Code:       "BAD_REQUEST",
		Message:    message,
		HTTPStatus: http.StatusBadRequest,
	}
}

func Validation(message string) error {
	return &Error{
		Code:       "VALIDATION_ERROR",
		Message:    message,
		HTTPStatus: http.StatusBadRequest,
	}
}

func Unauthorized(message string) error {
	return &Error{
		Code:       "UNAUTHORIZED",
		Message:    message,
		HTTPStatus: http.StatusUnauthorized,
	}
}

func Forbidden(message string) error {
	return &Error{
		Code:       "FORBIDDEN",
		Message:    message,
		HTTPStatus: http.StatusForbidden,
	}
}

func NotFound(message string) error {
	return &Error{
		Code:       "NOT_FOUND",
		Message:    message,
		HTTPStatus: http.StatusNotFound,
	}
}

func Conflict(message string) error {
	return &Error{
		Code:       "CONFLICT",
		Message:    message,
		HTTPStatus: http.StatusConflict,
	}
}

func Internal(message string) error {
	return &Error{
		Code:       "INTERNAL_SERVER_ERROR",
		Message:    message,
		HTTPStatus: http.StatusInternalServerError,
	}
}