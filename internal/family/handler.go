package family

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

func (h *Handler) CreateFamily(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var req CreateFamilyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
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
		switch {
		case errors.Is(err, ErrFamilyNotFound):
			response.NotFound(c, err.Error())
		default:
			response.Internal(c, "Internal server error")
		}
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

	if err := c.ShouldBindJSON(&req); err != nil {
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
		switch {
		case errors.Is(err, ErrOnlyOwnerCanInvite):
			response.Forbidden(c, err.Error())
		case errors.Is(err, ErrorAlreadyInvited):
			response.Conflict(c, err.Error())

		default: 
			response.Internal(c, "Internal server errror")
		}

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
		switch {
		case errors.Is(err, ErrInvitationNotFound): response.NotFound(c, err.Error())
		case errors.Is(err, ErrInvitationProcessed): response.BadRequest(c, err.Error())
		default : response.Internal(c, "Internal server error")
		}
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