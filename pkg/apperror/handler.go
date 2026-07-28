package apperror

import (
	"errors"

	"evara-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context, err error) {
	var appErr *Error

	if errors.As(err, &appErr) {
		response.Error(
			c,
			appErr.HTTPStatus,
			appErr.Code,
			appErr.Message,
		)
		return
	}

	response.Internal(c, "Internal server error")
}