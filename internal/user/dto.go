package user

type MeResponse struct {
	UserID string  `json:"user_id"`
	Name   *string `json:"name"`
	Email  string  `json:"email"`
}