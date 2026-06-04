package models

type CreateFamilyRequest struct {
	Name string `json:"name" binding:"required"`
}