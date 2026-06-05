package models

type CreateFamilyRequest struct {
	Name string `json:"name" binding:"required"`
}

type InvitationResponse struct {
	ID string `json:"id"`
	FamilyID string `json:"family_id"`
	FamilyName string `json:"family_name"`
	Status string `json:"status"` 
}