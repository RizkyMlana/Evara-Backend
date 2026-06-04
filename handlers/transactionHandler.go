package handlers

import (
	"context"
	"logqian-backend/config"
	"logqian-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



func CreateTransaction(c *gin.Context) {
	userID:= c.GetString("user_id")

	var input models.Transaction

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec(context.Background(),
	`insert into transactons (family_id, user_id, type, title, description, amount) values ($1, $2, $3, $4, $5, $6)`,
	input.FamilyID,
	userID,
	input.Type,
	input.Title,
	input.Description,
	input.Amount,
)

if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}


	c.JSON(http.StatusOK, gin.H {
		"message": "transaction created",
	})
}

func GetTransactions(c *gin.Context) {
	familyID := c.Query("family_id")
	c.JSON(http.StatusOK, gin.H{
		"message": "get transactions",
		"family_id": familyID,
		"data": []string{},
	})
}


