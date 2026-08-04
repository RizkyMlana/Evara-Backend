package family

import "time"

type Family struct {
	ID        string
	Name      string
	CreatedBy string
	CreatedAt time.Time
}

type FamilyMember struct {
	ID    string
	Name  *string
	Email string
	Role  string
}

type Invitation struct {
	ID        string
	FamilyID  string
	Email     string
	InvitedBy string
	Status    string
}