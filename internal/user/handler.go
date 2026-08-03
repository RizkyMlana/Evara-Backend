package user

import (
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

	data := h.service.Me(userID)
	response.OK(
		c,
		"success",
		data,
	)
}