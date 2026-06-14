package handlers

import (
	"context"
	"evara-backend/config"
	"evara-backend/models"

	"github.com/gin-gonic/gin"
)



func CreateTransaction(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	var input models.Transaction

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if input.Amount <= 0 {
		c.JSON(400, gin.H{
			"error": "amount must be greater than 0",
		})
		return
	}

	if input.Type != "income" && input.Type != "expense" {
		c.JSON(400, gin.H{
			"error": "invalid transaction type",
		})
		return
	}

	var exists bool

	err := config.DB.QueryRow(
		context.Background(),
		`
			select exists (
				select 1
				from family_members
				where family_id = $1
				and user_id = $2
			)
		`,
		input.FamilyID,
		userID,
	).Scan(&exists)

	if err != nil {
		c.JSON(500, gin.H{
			"error" : err.Error(),
		})
		return
	}

	if !exists {
		c.JSON(403, gin.H{
			"error": "not a family member",
		})
		return
	}

	var transactionID string

	err = config.DB.QueryRow(
		context.Background(),
		`
		insert into transactions (
			family_id,
			user_id,
			type,
			title,
			description,
			amount
		)
		values ($1, $2, $3, $4, $5, $6)
		returning id
		`,
		input.FamilyID,
		userID,
		input.Type,
		input.Title,
		input.Description,
		input.Amount,
	).Scan(&transactionID)

	if err != nil{
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "transaction created",
		"id": transactionID,
	})

}



func GetTransactions(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	familyID := c.Query("family_id")

	if familyID == "" {
		c.JSON(400, gin.H{
			"error": "famliy_id is required",
		})
		return
	}
	var isMember bool
	err := config.DB.QueryRow(
		context.Background(),
		`
			select exists(
				select 1
				from family_members
				where family_id = $1
				and user_id = $2
			)
		`,
		familyID,
		userID,
	).Scan(&isMember)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !isMember {
		c.JSON(403, gin.H{
			"error": "not a family member",
		})
		return
	}

	rows, err := config.DB.Query(
		context.Background(),
		`
			select t.id, coalesce(p.name, ''), t.type, t.title, t.description, t.amount, t.created_at
			from transactions t
			join profiles p on p.id = t.user_id
			where t.family_id = $1
			order by t.created_at desc
		`,
		familyID,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	transactions := []gin.H{}

	for rows.Next(){
		var id string
		var userName string
		var txType string
		var title string
		var description string
		var amount float64
		var createdAt string

		err := rows.Scan(
			&id, 
			&userName,
			&txType,
			&title,
			&description,
			&amount,
			&createdAt,
		)

		if err != nil {
			continue
		}

		transactions = append(transactions, gin.H{
			"id": id,
			"user_name": userName,
			"type": txType,
			"title": title,
			"description": description,
			"amount": amount,
			"created_at": createdAt,
		})
	}

	c.JSON(200, gin.H{
		"transactions": transactions,
	})
}


