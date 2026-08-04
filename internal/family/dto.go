package family

type CreateFamilyRequest struct {
	Name string `json:"name" validate:"required,min=3,max=50"`
}

type CreateFamilyResponse struct {
	ID string `json:"id"`
}

type GetMyFamilyResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type FamilyMemberResponse struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
	Role  string  `json:"role"`
}

type InviteMemberRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type InvitationResponse struct {
	ID         string `json:"id"`
	FamilyID   string `json:"family_id"`
	FamilyName string `json:"family_name"`
	Status     string `json:"status"`
}