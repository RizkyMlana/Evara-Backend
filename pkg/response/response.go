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
	Success(
		c, 
		http.StatusOK,
		message,
		data,
	)
}

func Created(
	c *gin.Context,
	message string,
	data any,
) {
	Success(
		c,
		http.StatusCreated,
		message,
		data,
	)
}

func Success(
	c *gin.Context,
	status int,
	message string,
	data any,
) {
	JSON(c, status, Response{
		Success: true,
		Message: message,
		Data: data,
	})
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func Paginated(
	c *gin.Context,
	message string,
	data any,
	pagination *Pagination,
){
	JSON(
		c,
		http.StatusOK,
		Response{
			Success: true,
			Message: message,
			Data: data,
			Pagination: pagination,
		},
	)
}