package family

import "errors"

var (
	ErrFamilyNotFound = errors.New("family not found")

	ErrOnlyOwnerCanInvite = errors.New("only owner can invite members")
	
	ErrorAlreadyInvited = errors.New("user already invited")

	ErrInvitationNotFound = errors.New("invitation not found")

	ErrInvitationProcessed = errors.New("invitation already processed")

)