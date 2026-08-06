package dashboard

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
	return &Handler{
		service: service,
	}
}

func (h *Handler)GetDashboard(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	familyID := c.Query("family_id")

	if familyID == "" {
		response.BadRequest(c, "family_id is required")
		return
	}

	data, err := h.service.GetDashboard(
		c.Request.Context(),
		userID,
		familyID,
	)

	if err != nil {
		apperror.Handle(c, err)
		return
	}
	response.OK(
		c,
		"dashboard fetched",
		data,
	)
}
