package dashboard

import (
	"context"
	"evara-backend/pkg/apperror"
	"time"
)

type Service interface {
	GetDashboard(
		ctx context.Context,
		userID string,
		familyID string,
	)(*DashboardResponse, error) 
}

type service struct {
	repository Repository
}

func NewService(
	repository Repository,
) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetDashboard(
	ctx context.Context,
	userID string,
	familyID string,
) (*DashboardResponse, error) {
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
		return nil, err
	}

	if !isMember {
		return nil, apperror.Forbidden("you are not a member of this family")
	}
	dashboard, err := s.repository.GetSummary(ctx, familyID)
	return &DashboardResponse{
		TotalIncome: dashboard.TotalIncome,
		TotalExpense: dashboard.TotalExpense,
		Balance: dashboard.TotalIncome - dashboard.TotalExpense,
	},nil
}