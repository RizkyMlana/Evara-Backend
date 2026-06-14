package models

type Transaction struct {
	FamilyID	string `json:"family_id" binding:"required"`
	Type	string 	`json:"type" binding:"required"`
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
	Amount float64 `json:"amount" binding:"required"`
}
