package family

import (
	"context"
	"errors"
	"net/http"

	"evara-backend/internal/config"
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

func GetFamilyMember (c *gin.Context) {
	familyID := c.Param("id")


	rows, err := config.DB.Query(
		context.Background(),
		`
			select p.id, p.name, p.email, fm.role
			from family_members fm
			join profiles p on p.id = fm.user_id
			where fm.family_id = $1	
		`,
		familyID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()
	members := []gin.H{}

	for rows.Next() { 
		var id string
		var name *string
		var email string
		var role string

		err := rows.Scan(
			&id,
			&name,
			&email,
			&role,
		)

		if err != nil {
			continue
		}


		members = append(members, gin.H{
			"id": id,
			"name": name,
			"email": email,
			"role": role,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"members": members,
	})
}

func InviteMember (c *gin.Context) {
	userID, _ := c.Get("user_id")
	familyID := c.Param("id")


	var input struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var role string

	err := config.DB.QueryRow(
		context.Background(),
		`select role
		 from family_members
		 where family_id = $1
		 and user_id = $2
		`,
		familyID,
		userID,
	).Scan(&role)



	if err != nil || role != "owner"{
		c.JSON(403, gin.H{
			"error": "only owner can invite members",
		})
		return
	}

	_, err = config.DB.Exec(
		context.Background(),	
		`
			insert into invitations (family_id, email, invited_by)
			values ($1, $2, $3)
		`,
		familyID,
		input.Email,
		userID,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

	}
}

func GetInvitations (c *gin.Context) {
	userID, _ := c.Get("user_id")

	rows, err := config.DB.Query(
		context.Background(),
		`
			select i.id, i.family_id, f.name, i.status
			from invitations i
			join families f on f.id = i.family_id
			join profiles p on p.email = i.email
			where p.id = $1
			and i.status = 'pending'	
		`,
		userID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()
	var invitations []InvitationResponse

	for rows.Next() {
		var invite InvitationResponse

		err := rows.Scan(
			&invite.ID,
			&invite.FamilyID,
			&invite.FamilyName,
			&invite.Status,
		)

		if err != nil {
			continue
		}

		invitations = append(invitations, invite)

	}
	c.JSON(http.StatusOK, gin.H{
		"invitations": invitations,
	})
}

func AcceptInvitations (c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	invitationID := c.Param("id")

	var familyID string 
	var status string

	err := config.DB.QueryRow(
		context.Background(),
		`
			select family_id, status
			from invitations
			where id = $1 
		`,
		invitationID,
	).Scan(
		&familyID,
		&status,
	)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "invitation not found",
		})
		return
	}

	if status != "pending" {
		c.JSON(400, gin.H{
			"error": "invitation already processed",
		})
		return
	}
	_, err = config.DB.Exec(
		context.Background(),
		`
			insert into family_members (family_id, user_id, role)
			values ($1, $2, 'member')
		`,
		familyID,
		userID,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = config.DB.Exec(
		context.Background(),
		`
			update invitations
			set status = 'accepted'
			where id = $1
		`,
		invitationID,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "invitation accepted",
	})
}

func RejectInvitation(c *gin.Context) {
	invitationID := c.Param("id")

	_, err := config.DB.Exec(
		context.Background(),
		`
			update invitations
			set status = 'rejected'
			where id = $1
			and status = 'pending'
		`,
		invitationID,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "invitation rejected",
	})
}

func (h *Handler) GetFamilyMember(c *gin.Context) {
	familyID := c.Param("id")
	members, err := h.service.GetMembers(
		c.Request.Context(),
		familyID,
	)

	if err != nil {
		response.Internal(
			c,
			"Internal server error",
		)
		return
	}

	response.OK(
		c,
		"Family members retrieved successfully",
		members,
	)
}