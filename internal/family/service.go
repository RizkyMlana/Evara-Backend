package family

import "context"

type Service interface {
	Create(
		ctx context.Context,
		userID string,
		req CreateFamilyRequest,
	) (string, error)

	GetMyFamily(
		ctx context.Context,
		userID string,
	) (*GetMyFamilyResponse, error)

	GetMembers(
		ctx context.Context,
		familyID string,
	)([]FamilyMemberResponse, error)

}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(
	ctx context.Context,
	userID string,
	req CreateFamilyRequest,
) (string, error) {

	familyID, err := s.repository.Create(
		ctx,
		userID,
		req.Name,
	)
	if err != nil {
		return "", err
	}

	err = s.repository.AddOwner(
		ctx,
		familyID,
		userID,
	)
	if err != nil {
		return "", err
	}

	return familyID, nil
}

func (s *service) GetMyFamily(
	ctx context.Context,
	userID string,
) (*GetMyFamilyResponse, error) {
	return s.repository.GetMyFamily(
		ctx,
		userID,
	)
}

func (s *service)GetMembers(
	ctx context.Context,
	familyID string,
)([]FamilyMemberResponse, error){
	return s.repository.GetMembers(
		ctx,
		familyID,
	)
}