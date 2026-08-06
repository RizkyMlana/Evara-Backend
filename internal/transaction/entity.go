package transaction

import "time"

type Transaction struct {
	ID 					string
	FamilyID    string  	
	UserID 			string 		
	Type        string
	Title       string
	Description string  	
	Amount      float64 	
	CreatedAt 	time.Time	
}
