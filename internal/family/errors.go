package family

import "errors"

var (
	ErrFamilyNotFound = errors.New("family not found")

	ErrOnlyOwnerCanInvite =
		errors.New("only owner can invite members")

	ErrInvitationNotFound =
		errors.New("invitation not found")

	ErrInvitationProcessed =
		errors.New("invitation already processed")
)