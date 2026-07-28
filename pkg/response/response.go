package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success    bool         `json:"success"`
	Message    string       `json:"message,omitempty"`
	Data       any          `json:"data,omitempty"`
	Error      *ErrorDetail `json:"error,omitempty"`
	Pagination *Pagination  `json:"pagination,omitempty"`
}

func JSON(
	c *gin.Context,
	status int,
	res Response,
) {
	c.JSON(status, res)
}

func OK(
	c *gin.Context,
	message string,
	data any,
) {
	JSON(c, http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Created(
	c *gin.Context,
	message string,
	data any,
) {
	JSON(c, http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}