package dashboard

import "context"

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
	famiilyID string,
) (*DashboardResponse, error) {
	isMember, err := s.repository.IsFamilyMember(
		ctx,
		famiilyID,
		userID,
	)

	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, ErrNotFamilyMember
	}

	return s.repository.GetSummary(
		ctx,
		famiilyID,
	)
}