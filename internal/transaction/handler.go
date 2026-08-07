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


// CreateTransaction godoc
// @Summary Create transaction
// @Description Create a new income or expense transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateTransactionRequest true "Transaction data"
// @Success 201 {object} response.Response{data=CreateTransactionResponse}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /transaction [post]
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

// GetTransactions godoc
// @Summary Get transactions
// @Description Get all transactions form a family
// @Tags Transaction
// @Produce json
// @Security BearerAuth
// @Param family_id query string true "Family ID"
// @Success 200 {object} response.Response{data=[]GetTransactionsResponse}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /transaction [get]
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

// DeleteTransaction godoc
// @Summary Delete transaction
// @Description Delete a transaction by its ID
// @Tags Transaction
// @Produce json
// @Security BearerAuth
// @Param id path string true "Transaction ID"
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /transaction/{id} [delete]
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