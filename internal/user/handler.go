package user

import (
	"evara-backend/pkg/apperror"
	"evara-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(
	service Service,
) *Handler {
	return  &Handler{
		service: service,
	}
}

func (h *Handler) Me(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	data, err := h.service.Me(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		apperror.Handle(c, err)
		return
	}

	response.OK(
		c,
		"profiles retrieved successfully",
		data,
	)
}