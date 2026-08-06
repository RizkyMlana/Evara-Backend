package transaction

import (
	"context"
	"errors"
	"evara-backend/pkg/apperror"
	"fmt"
	"time"
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

	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()

	if req.Amount <= 0 {
		return "", apperror.BadRequest(
			"amount must be greater than zero",
		)
	}

	if req.Type != "income" &&
		req.Type != "expense" {
			return "", apperror.BadRequest(
				"invalid transaction type",
			)
	}

	isMember, err := s.repository.IsFamilyMember(
		ctx,
		req.FamilyID,
		userID,
	)
	if err != nil {
		return "", fmt.Errorf("check family member: %w", err)
	}
	if !isMember {
		return "", apperror.Forbidden(
			"you are not a member of this family",
		)
	}



	tx := Transaction{
		FamilyID: req.FamilyID,
		UserID: userID,
		Type: req.Type,
		Title: req.Title,
		Description: req.Description,
		Amount: req.Amount,
	}

	id, err := s.repository.Create(
		ctx, 
		tx,
	)
	if err != nil {
		return "", fmt.Errorf("create transaction: %w", err)
	}
	return id, nil
}

func (s *Service) GetTransactions(
	ctx context.Context,
	userID string,
	familyID string,
) ([]GetTransactionsResponse, error) {

	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()
	isMember, err := s.repository.IsFamilyMember(
		ctx,
		familyID,
		userID,
	)
	if err != nil {
		return  nil, fmt.Errorf("check family member: %w", err)
	}

	if !isMember {
		return  nil, apperror.Forbidden(
			"you are not a member of this family",
		)
	}

	transactions, err := s.repository.GetTransactions(
		ctx,
		familyID,
	)

	if err != nil{
		return nil, fmt.Errorf("get transactions: %w", err)
	}

	return  transactions, nil
}

func (s *Service) Delete(
	ctx context.Context,
	userID string,
	transactionID string,
) error {

	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()
	
	tx, err := s.repository.GetByID(
		ctx,
		transactionID,
	)
	
	if err != nil {
		if errors.Is(err, ErrTransactionNotFound) {
			return apperror.NotFound("transaction not found")
		}
	}

	if tx.UserID != userID {
		return apperror.Forbidden(
			"you can only delete your own transaction",
		)
	}

	err = s.repository.Delete(
		ctx,
		transactionID,
	)
	if err != nil {
		return fmt.Errorf("delete transaction: %w", err)
	}
	return nil

}