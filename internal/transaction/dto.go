package transaction

import "time"

type CreateTransactionRequest struct {
	FamilyID    string  `json:"family_id" validate:"required,uuid"`
	Type        string  `json:"type" validate:"required,oneof=income expense"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount" validate:"gt=0"`
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
	CreatedAt time.Time `json:"created_at"`
}
