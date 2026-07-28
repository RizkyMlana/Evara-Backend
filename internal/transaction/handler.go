package transaction

import (
	"errors"
	"evara-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateTransaction(c *gin.Context) {

	userID := c.MustGet("user_id").(string)

	var req CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Validation(c, err.Error())
	}

	id, err := h.service.Create(
		c.Request.Context(),
		userID,
		req,
	)

	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidAmount):
			response.BadRequest(c, err.Error())
		case errors.Is(err, ErrInvalidTransactionType):
			response.BadRequest(c, err.Error())
		case errors.Is(err, ErrNotFamilyMember):
			response.Forbidden(c, err.Error())
		default:
			response.Internal(c, "internal server error")
		}
		return
		

	}

	response.Created(
		c,
		"Transaction created succesfully",
		CreateTransactionResponse{
			ID: id,
		},
	)
}


func (h *Handler) GetTransactions(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	familyID := c.Query("family_id")

	if familyID == "" {
		response.Validation(
			c,
			"family_id is required",
		)
		return
	}

	transactions, err := h.service.GetTransactions(
		c.Request.Context(),
		userID,
		familyID,
	)

	if err != nil{
		switch {
		case errors.Is(err, ErrNotFamilyMember):
			response.Forbidden(
				c,
				err.Error(),
			)

		default:
			response.Internal(
				c,
				"Internal server error",
			)
		}
		return
	}
	response.OK(
		c,
		"Transaction retrieved succesfully",
		transactions,
	)

}

func (h *Handler) DeleteTransaction(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	transactionID := c.Param("id")

	err := h.service.Delete(
		c.Request.Context(),
		userID,
		transactionID,
	)

	if err != nil{
		switch {
		case errors.Is(err, ErrTransactionNotFound):
			response.NotFound(
				c,
				err.Error(),
			)
		case errors.Is(err, ErrForbiddenDelete):
			response.Forbidden(
				c,
				err.Error(),
			)
		default:
			response.Internal(
				c,
				"Internal server error",
			)
		}
		return
	}

	response.OK(
		c,
		"Transaction deleted successfully",
		nil,
	)
}