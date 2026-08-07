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

// GetMyFamily godoc
// @Summary Get my family
// @Description Get the authenticated user's family information
// @Tags Family
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=GetMyFamilyResponse}
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /families/me [get]
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

// GetFamilyMembers godoc
// @Summary Get family members
// @Description Get all members of a family
// @Tags Family
// @Produce json
// @Security BearerAuth
// @Param id path string true "Family ID"
// @Success 200 {object} response.Response{data=[]FamilyMemberResponse}
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /families/{id}/members [get]
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

// InviteMember godoc
// @Summary Invite family member
// @Description Invite a user to join a family by email
// @Tags Family
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Family ID"
// @Param request body InviteMemberRequest true "Invitation request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 409 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /families/{id}/invite [post]
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

// GetInvitations godoc
// @Summary Get invitations
// @Description Get all pending family invitations for the authenticated user
// @Tags Family
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]InvitationResponse}
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /families/invitations [get]
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

// AcceptInvitation godoc
// @Summary Accept invitation
// @Description Accept a pending family invitation
// @Tags Family
// @Produce json
// @Security BearerAuth
// @Param id path string true "Invitation ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /families/invitations/{id}/accept [patch]
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

// RejectInvitation godoc
// @Summary Reject invitation
// @Description Reject a pending family invitation
// @Tags Family
// @Produce json
// @Security BearerAuth
// @Param id path string true "Invitation ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /families/invitations/{id}/reject [patch]
func (h *Handler) RejectInvitation(c *gin.Context) {
	invitationID := c.Param("id")

	err := h.service.RejectInvitation(
		c.Request.Context(),
		invitationID,
	)

	if err != nil {
		apperror.Handle(c, err)
		return
	}

	response.OK(
		c,
		"Invitation rejected successfully",
		nil,
	)
}