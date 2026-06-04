package models

type Transaction struct {
	FamilyID	string `json:"family_id"`
	Type	string 	`json:"type"`
	Title string `json:"title"`
	Description string `json:"description"`
	Amount float64 `json:"amount"`
}