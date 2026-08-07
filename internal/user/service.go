package user

import (
	"context"
	"fmt"
	"time"
)

type Service interface {
	Me(
		ctx context.Context,
		userID string,
		) (*MeResponse, error)
}

type service struct {
	repository Repository
}

func NewService (
	repository Repository,
) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Me(
	ctx context.Context,
	userID string,
) (*MeResponse, error) {
	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()
	user, err := s.repository.GetByID(
		ctx,
		userID,
	)

	if err != nil {
		fmt.Printf("GetByID error: %T %+v\n", err, err)
		return nil, err
	}
	fmt.Printf("User: %+v\n", user)

	return &MeResponse{
		UserID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}, nil
}