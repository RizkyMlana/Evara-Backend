package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Error(
	c *gin.Context,
	status int,
	code string,
	message string,
) {
	JSON(c, status, Response{
		Success: false,
		Error: &ErrorDetail{
			Code:    code,
			Message: message,
		},
	})
}

func BadRequest(c *gin.Context, err string) {
	Error(c, http.StatusBadRequest, CodeBadRequest, err)
}

func Unauthorized(c *gin.Context, err string) {
	Error(c, http.StatusUnauthorized, CodeUnauthorized, err)
}

func Forbidden(c *gin.Context, err string) {
	Error(c, http.StatusForbidden, CodeForbidden, err)
}

func NotFound(c *gin.Context, err string) {
	Error(c, http.StatusNotFound, CodeNotFound, err)
}

func Conflict(c *gin.Context, err string) {
	Error(c, http.StatusConflict, CodeConflict, err)
}

func Validation(c *gin.Context, err string) {
	Error(c, http.StatusBadRequest, CodeValidation, err)
}

func Internal(c *gin.Context, err string) {
	Error(c, http.StatusInternalServerError, CodeInternal, err)
}
