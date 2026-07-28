package transaction

type CreateTransactionRequest struct {
	FamilyID    string  `json:"family_id" binding:"required"`
	Type        string  `json:"type" binding:"required"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount" binding:"required"`
}

type TransactionResponse struct {
	ID          string  `json:"id"`
	UserName    string  `json:"user_name"`
	Type        string  `json:"type"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	CreatedAt   string  `json:"created_at"`
}

type CreateTransactionResponse struct {
	ID string `json:"id"`
}

type GetTransactionsResponse struct {
	ID string `json:"id"`
	UserName string `json:"user_name"`
	Type string `json:"type"`
	Title string `json:"title"`
	Description string `json:"description"`
	Amount float64 `json:"amount"`
	CreatedAt string `json:"created_at"`
}
