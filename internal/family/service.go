package family

import (
	"context"
	"evara-backend/internal/database"
	"evara-backend/pkg/apperror"
	"time"
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
	txManager *database.TransactionManager
}

func NewService(
	repository Repository, 
	txManager *database.TransactionManager,
	) Service {
	return &service{
		repository: repository,
		txManager: txManager,
	}
}

func (s *service) Create(
	ctx context.Context,
	userID string,
	req CreateFamilyRequest,
) (string, error) {

	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()
	var familyID string

	err := s.txManager.WithinTransaction(
		ctx,
		func (tx database.DBTX) error {
			repo := NewRepository(tx)

			id, err := repo.Create(
				ctx,
				userID,
				req.Name,
			)
			if err != nil {
				return err
			}
			if err := repo.AddOwner(
				ctx,
				id,
				userID,
			); err != nil {
				return err
			}
			familyID = id
			return nil
		},
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

	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()

	return s.repository.GetMyFamily(
		ctx,
		userID,
	)
}

func (s *service)GetMembers(
	ctx context.Context,
	familyID string,
)([]FamilyMemberResponse, error){

	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()
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

	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()
	role, err := s.repository.GetRole(
		ctx,
		familyID,
		userID,
	)

	if err != nil {
		return err
	}

	if role != "owner" {
		return apperror.Forbidden(
			"only owner can invite member",
		)
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
		return apperror.Conflict(
			"user already invited",
		)
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
	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()
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
	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()
	return s.txManager.WithinTransaction(
		ctx,
		func(tx database.DBTX) error {
			repo := NewRepository(tx)

			invitation, err := repo.GetInvitation(
				ctx,
				invitationID,
			)

			if err != nil {
				return err
			}
			if invitation.Status != "pending" {
				return apperror.BadRequest(
					"invitation already processed",
				)
			}

			err = repo.AddMember(
				ctx,
				invitation.FamilyID,
				userID,
			)

			if err != nil {
				return err
			}
			err = repo.UpdateInvitationStatus(
				ctx,
				invitationID,
				"accepted",
			)
			if err != nil {
				return err
			}
			return nil
		},
	)

}

func (s *service)RejectInvitation(
	ctx context.Context,
	invitationID string,
) error {
	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()
	return s.repository.RejectInvitation(
		ctx,
		invitationID,
	)
}