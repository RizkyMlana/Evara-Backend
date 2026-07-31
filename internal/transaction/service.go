package transaction

import (
	"context"
)

type Service struct {
	repository Repository
}

func NewService(
	repository Repository,
) *Service {

	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(
	ctx context.Context,
	userID string,
	req CreateTransactionRequest,
) (string, error) {

	if req.Amount <= 0 {
		return "", ErrInvalidAmount
	}

	if req.Type != "income" &&
		req.Type != "expense" {
			return "", ErrInvalidTransactionType
	}

	isMember, err := s.repository.IsFamilyMember(
		ctx,
		req.FamilyID,
		userID,
	)
	if err != nil {
		return "", err
	}
	if !isMember {
		return "", ErrNotFamilyMember
	}



	tx := Transaction{
		FamilyID: req.FamilyID,
		UserID: userID,
		Type: req.Type,
		Title: req.Title,
		Description: req.Description,
		Amount: req.Amount,
	}

	return s.repository.Create(
		ctx,
		tx,
	)
}

func (s *Service) GetTransactions(
	ctx context.Context,
	userID string,
	familyID string,
) ([]GetTransactionsResponse, error) {
	isMember, err := s.repository.IsFamilyMember(
		ctx,
		familyID,
		userID,
	)
	if err != nil {
		return  nil, err
	}

	if !isMember {
		return  nil, ErrNotFamilyMember
	}

	return s.repository.GetTransactions(
		ctx,
		familyID,
	)
}

func (s *Service) Delete(
	ctx context.Context,
	userID string,
	transactionID string,
) error {
	tx, err := s.repository.GetByID(
		ctx,
		transactionID,
	)
	if err != nil {
		return  err
	}

	if tx.UserID != userID {
		return ErrForbiddenDelete
	}

	return s.repository.Delete(
		ctx,
		transactionID,
	)

}