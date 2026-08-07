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

// GetDashboard godoc
// @Summary Get dashboard
// @Description Get dashboard summary for a family
// @Tags Dashboard
// @Produce json
// @Security BearerAuth
// @Param family_id query string true "Family ID"
// @Success 200 {object} response.Response{data=DashboardResponse} 
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /dashboard [get]
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
