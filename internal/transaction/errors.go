package transaction

import "errors"

var (
	ErrInvalidAmount = errors.New("amount must be greater than zero")
	ErrInvalidTransactionType = errors.New("invalid transaction type")
	ErrNotFamilyMember = errors.New("user is not a family member")
	ErrTransactionNotFound    = errors.New("transaction not found")
	ErrForbiddenDelete = errors.New("you can only delete your own transaction")
)