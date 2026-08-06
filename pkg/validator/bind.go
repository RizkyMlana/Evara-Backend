package validator

import (
	"evara-backend/pkg/apperror"

	"github.com/gin-gonic/gin"
)

func Bind(c *gin.Context, v any) error {
	if err := c.ShouldBindJSON(v); err != nil {
		return apperror.Validation(err.Error())
	}

	if err := Validate(v); err != nil {
		return apperror.Validation(err.Error())
	}
	return nil
}