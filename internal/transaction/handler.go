package transaction

import (
	"evara-backend/pkg/apperror"
	"evara-backend/pkg/response"
	"evara-backend/pkg/validator"

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
		apperror.Handle(c, err)
		return
	}

	response.Created(
		c,
		"Transaction created successfully",
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
		apperror.Handle(c, err)
		return
	}
	response.OK(
		c,
		"Transaction retrieved successfully",
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
		apperror.Handle(c, err)
		return
	}

	response.OK(
		c,
		"Transaction deleted successfully",
		nil,
	)
}