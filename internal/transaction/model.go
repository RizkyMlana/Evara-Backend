package transaction

import "time"

type Transaction struct {
	ID 					string 		`json:"id"`
	FamilyID    string  	`json:"family_id" binding:"required"`
	UserID 			string 		`json:"user_id"`
	Type        string  	`json:"type" binding:"required"`
	Title       string  	`json:"title" binding:"required"`
	Description string  	`json:"description"`
	Amount      float64 	`json:"amount" binding:"required"`
	CreatedAt 	time.Time	`json:"created_at"`
}
