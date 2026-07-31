package family

import (
	"context"
)

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

	Invite(
		ctx context.Context,
		userID string,
		familyID string,
		req InviteMemberRequest,
	) error
	
	GetInvitations(
		ctx context.Context,
		userID string,
	) ([]InvitationResponse, error)

	AcceptInvitation(
		ctx context.Context,
		userID string,
		invitationID string,
	) error

	RejectInvitation(
		ctx context.Context,
		invitationID string,
	) error
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

func(s *service)Invite(
	ctx context.Context,
	userID string,
	familyID string,
	req InviteMemberRequest,
) error {
	role, err := s.repository.GetRole(
		ctx,
		familyID,
		userID,
	)

	if err != nil {
		return err
	}

	if role != "owner" {
		return ErrOnlyOwnerCanInvite
	}

	exists, err := s.repository.InvitationExists(
		ctx,
		familyID,
		req.Email,
	)

	if err != nil {
		return err
	}

	if exists {
		return ErrorAlreadyInvited
	}

	return s.repository.CreateInvitation(
		ctx,
		familyID,
		req.Email,
		userID,
	)
}


func (s *service) GetInvitations(
	ctx context.Context,
	userID string,
)([]InvitationResponse, error) {
	return s.repository.GetInvitations(
		ctx,
		userID,
	)
}
func (s *service) AcceptInvitation(
	ctx context.Context,
	userID string,
	invitationID string,
) error {
	invitation, err := s.repository.GetInvitation(
		ctx,
		invitationID,
	)

	if err != nil {
		return err
	}

	if invitation.Status != "pending" {
		return ErrInvitationProcessed
	}

	err = s.repository.AddMember(
		ctx,
		invitation.FamilyID,
		userID,
	)

	if err != nil {
		return err
	}

	return s.repository.UpdateInvitationStatus(
		ctx,
		invitationID,
		"accepted",
	)
}

func (s *service)RejectInvitation(
	ctx context.Context,
	invitationID string,
) error {
	return s.repository.RejectInvitation(
		ctx,
		invitationID,
	)
}