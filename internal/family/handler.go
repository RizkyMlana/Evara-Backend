package family

import (
	"evara-backend/pkg/apperror"
	"evara-backend/pkg/response"
	"evara-backend/pkg/validator"

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


// CreateFamily godoc
// 
// @Summary Create family
// @Description Create a new family
// @Tags Family
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body CreateFamilyRequest true "Create Family"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /families [post]
func (h *Handler) CreateFamily(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var req CreateFamilyRequest

	if err := validator.Bind(c, &req); err != nil {
		response.Validation(c, err.Error())
		return
	}

	id, err := h.service.Create(
		c.Request.Context(),
		userID,
		req,
	)

	if err != nil {
		response.Internal(c, err.Error())
		return
	}

	response.Created(
		c,
		"Family created successfully",
		CreateFamilyResponse{
			ID: id,
		},
	)
}

func (h *Handler) GetMyFamily(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	family, err := h.service.GetMyFamily(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		apperror.Handle(c, err)
		return
	}

	response.OK(
		c,
		"Family retrieved successfully",
		family,
	)
}

func(h *Handler) GetFamilyMember(c *gin.Context) {
	familyID := c.Param("id")
	members, err := h.service.GetMembers(
		c.Request.Context(),
		familyID,
	)

	if err != nil {
		response.Internal(c, "Internal server error")
		return
	}
	response.OK(
		c,
		"Family members retrieved successfully",
		members,
	)
}

func (h *Handler)InviteMember(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	familyID := c.Param("id")

	var req InviteMemberRequest

	if err := validator.Bind(c, &req); err != nil {
		response.Validation(c, err.Error())
		return
	}

	err := h.service.Invite(
		c.Request.Context(),
		userID,
		familyID,
		req,
	)

	if err != nil {
		apperror.Handle(c, err)
		return
	}

	response.OK(
		c,
		"Invitation sent successfully",
		nil,
	)
}

func (h *Handler)GetInvitations(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	invitations, err := h.service.GetInvitations(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		response.Internal(c, "Internal server error")
		return
	}

	response.OK(
		c,
		"Invitations retrieved successfully",
		invitations,
	)
}

func (h *Handler) AcceptInvitation(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	invitationID := c.Param("id")

	err := h.service.AcceptInvitation(
		c.Request.Context(),
		userID,
		invitationID,
	)
	if err != nil {
		apperror.Handle(c, err)
		return
	}

	response.OK(
		c,
		"Invitation accepted successfully",
		nil,
	)
}

func (h *Handler) RejectInvitation(c *gin.Context) {
	invitationID := c.Param("id")

	err := h.service.RejectInvitation(
		c.Request.Context(),
		invitationID,
	)

	if err != nil {
		response.Internal(c, "Internal server error")
		return
	}

	response.OK(
		c,
		"Invitation rejected successfully",
		nil,
	)
}