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


// Me godoc
// @Summary Get my profile
// @Description Get the authenticated user's profile information
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {objec} response.Response{data=MeResponse}
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /me [get]
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