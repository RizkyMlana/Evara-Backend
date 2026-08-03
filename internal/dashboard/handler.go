package dashboard

import (
	"errors"
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
		switch {
		case errors.Is(err, ErrNotFamilyMember):
			response.Forbidden(c, err.Error())

		default:
			response.Internal(c, err.Error())
		}
		return
	}
	response.OK(
		c,
		"dashboard fetched",
		data,
	)
}
