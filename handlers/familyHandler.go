package handlers

import (
	"context"
	"net/http"

	"logqian-backend/config"
	"logqian-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateFamily (c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input models.CreateFamilyRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),

		})
		return
	}

	var familyID string

	err := config.DB.QueryRow(
		context.Background(),
		`
			insert into families (name, created_by)
			values ($1, $2) returning id
		`,
		input.Name,
		userID,
	).Scan(&familyID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	_, err = config.DB.Exec(
		context.Background(),
		`insert into family_members (family_id, user_id, role)
		 values ($1, $2, 'owner')
		`,
		familyID,
		userID,
	)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message" : "familty created",
		"family_id" : familyID,
	})
}

func GetMyFamily (c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	var familyID string
	var familyName string
	var role string

	err := config.DB.QueryRow(
		context.Background(),
		`
			select f.id, f.name, fm.role
			from family_members fm
			join families f on f.id = fm.family_id
			where fm.user_id = $1
			limit 1
		`,
		userID,
	).Scan(
		&familyID,
		&familyName,
		&role,
	)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "family not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"id": familyID,
		"name" : familyName,
		"role" : role,
	})
}