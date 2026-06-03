package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	FamilyID	string `json:"family_id"`
	Type	string 	`json:"type"`
	Title string `json:"title"`
	Description string `json:"description"`
	Amount float64 `json:"amount"`
}

func CreateTransaction(c *gin.Context) {
	userID:= c.GetString("user_id")

	var input Transaction

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"message": "transaction created",
		"user_id": userID,
		"data": input,
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