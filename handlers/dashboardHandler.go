package handlers

import (
	"context"
	"evara-backend/config"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	familyID := c.Query("family_id")

	if familyID == "" {
		c.JSON(400, gin.H{
			"error": "family_id is required",
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

	var totalIncome float64
	var totalExpense float64
	
	err = config.DB.QueryRow(
		context.Background(),
		`
			select coalesce(sum(amount),0)
			from transactions
			where family_id = $1
			and type = 'income'
		`,
		familyID,
	).Scan(&totalIncome)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = config.DB.QueryRow(
		context.Background(),
		`
			select coalesce(sum(amount),0)
			from transactions
			where family_id = $1
			and type = 'expense'
		`,
		familyID,
	).Scan(&totalExpense)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"total_income": totalIncome,
		"total_expense": totalExpense,
		"balance": totalIncome - totalExpense,
	})
}