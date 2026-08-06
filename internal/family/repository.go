package family

import (
	"context"
	"errors"
	"evara-backend/internal/database"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Create(
		ctx context.Context,
		userID string,
		name string,
	) (string, error)

	AddOwner(
		ctx context.Context,
		familyID string,
		userID string,
	) error

	GetMyFamily(
		ctx context.Context,
		userID string,
	) (*GetMyFamilyResponse, error)

	GetMembers(
		ctx context.Context,
		familyID string,
	) ([]FamilyMemberResponse, error)

	GetRole(
		ctx context.Context,
		familyID string,
		userID string,
	) (string, error)

	InvitationExists(
		ctx context.Context,
		familyID string,
		userID string,
	) (bool, error)

	CreateInvitation(
		ctx context.Context,
		familyID string,
		email string,
		userID string,
	) error
	
	GetInvitations(
		ctx context.Context,
		userID string,
	) ([]InvitationResponse, error)

	GetInvitation(
		ctx context.Context,
		id string,
	)(*Invitation, error)
	
	AddMember(
		ctx context.Context,
		familyID string,
		userID string, 
	) error

	UpdateInvitationStatus(
		ctx context.Context,
		id string,
		status string,
	) error

	RejectInvitation(
		ctx context.Context,
		id string,
	) error

}

type repository struct{
	db database.DBTX
}

func NewRepository(
	db database.DBTX,
) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(
	ctx context.Context,
	userID string,
	name string,
) (string, error) {

	var familyID string

	err := r.db.QueryRow(
		ctx,
		`
		INSERT INTO families (name, created_by)
		VALUES ($1, $2)
		RETURNING id
		`,
		name,
		userID,
	).Scan(&familyID)

	if err != nil {
		return "", fmt.Errorf("create family: %w", err)
	}

	return familyID, nil
}

func (r *repository) AddOwner(
	ctx context.Context,
	familyID string,
	userID string,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO family_members (
			family_id,
			user_id,
			role
		)
		VALUES ($1, $2, 'owner')
		`,
		familyID,
		userID,
	)
	if err != nil {
		return fmt.Errorf("add family owner: %w", err)
	}
	return nil
}

func (r *repository) GetMyFamily(
	ctx context.Context,
	userID string,
) (*GetMyFamilyResponse, error) {
	var family GetMyFamilyResponse

	err := r.db.QueryRow(
		ctx,
		`
			select f.id, f.name, fm.role
			from family_members fm
			join families f on f.id = fm.family_id
			where fm.user_id = $1
			limit 1	
		`,
		userID,
	).Scan(
		&family.ID,
		&family.Name,
		&family.Role,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrFamilyNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get my family: %w", err)
	}

	return &family, nil
}

func (r *repository) GetMembers(
	ctx context.Context,
	familyID string,
)([]FamilyMemberResponse, error){
	rows, err := r.db.Query(
		ctx,
		`
			select p.id, p.name, p.email, fm.role
			from family_members fm
			join profiles p on p.id = fm.user_id
			where fm.family_id = $1	
		`,
		familyID,
	)
	if err != nil {
		return nil, fmt.Errorf("get family members: %w", err)
	}

	defer rows.Close()

	var members []FamilyMemberResponse
	for rows.Next() {
		var member FamilyMemberResponse

		err := rows.Scan(
			&member.ID,
			&member.Name,
			&member.Email,
			&member.Role,
		)

		if err != nil {
			return nil, fmt.Errorf("scan family member: %w", err)
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate family members: %w", err)
	}

	return members, nil

}

func (r *repository)GetRole(
	ctx context.Context,
	familyID string,
	userID string,
) (string, error) {
	var role string

	err := r.db.QueryRow(
		ctx,
		`
			select role
			from family_members
			where family_id = $1
			and user_id = $2	
		`,
		familyID,
		userID,
	).Scan(&role)

	if err != nil {
		return "", fmt.Errorf("get member role: %w", err)
	}

	return role, nil
}

func (r * repository) InvitationExists(
	ctx context.Context,
	familyID string,
	email string,
) (bool, error) {
	var exists bool

	err := r.db.QueryRow(
		ctx,
		`
			select exists(
				select 1
				from invitations
				where family_id = $1
				and email = $2
				and status = 'pending'
			)	
		`,
		familyID,
		email,
	).Scan(&exists)
	
	if err != nil {
		return false, fmt.Errorf("check invitation exists: %w", err)
	}
	return exists, nil
}

func (r *repository) CreateInvitation(
	ctx context.Context,
	familyID string,
	email string,
	userID string,
) error {
	_, err := r.db.Exec(
		ctx,
		`
			insert into invitations(
				family_id,
				email,
				invited_by
			)	
			values($1, $2, $3)
		`,
		familyID,
		email,
		userID,
	)
	if err != nil {
		return fmt.Errorf("create invitation: %w", err)
	}
	return  nil
}

func (r *repository)GetInvitations(
	ctx context.Context,
	userID string,
)([]InvitationResponse, error) {
	rows, err := r.db.Query(
		ctx,
		`
			select i.id, i.family_id, f.name, i.status
			from invitations i
			join families f on f.id = i.family_id
			join profiles p on p.email = i.email
			where p.id = $1
			and i.status = 'pending'	
		`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("get invitations: %w", err)
	}

	defer rows.Close()

	var invitations []InvitationResponse

	for rows.Next() {
		var invitation InvitationResponse

		err := rows.Scan(
			&invitation.ID,
			&invitation.FamilyID,
			&invitation.FamilyName,
			&invitation.Status,
		)

		if err != nil {
			return nil, fmt.Errorf("scan invitation: %w", err)
		}

		invitations = append(invitations, invitation)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate invitations: %w", err)
	}

	return invitations, nil
}

func (r *repository)GetInvitation(
	ctx context.Context,
	id string,
)(*Invitation, error) {
	var invitation Invitation

	err := r.db.QueryRow(
		ctx,
		`
			select id, family_id, email, invited_by, status
			from invitations
			where id = $1	
		`,
		id,
	).Scan(
		&invitation.ID,
		&invitation.FamilyID,
		&invitation.Email,
		&invitation.InvitedBy,
		&invitation.Status,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrInvitationNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("get invitation: %w", err)
	}

	return &invitation, nil
}

func (r *repository) AddMember (
	ctx context.Context,
	familyID string,
	userID string,
) error {
	_, err := r.db.Exec(
		ctx,
		`
			insert into family_members(family_id, user_id, role)
			values($1, $2, 'member')	
		`,
		familyID,
		userID,
	)
	if err != nil {
		return fmt.Errorf("add family member: %w", err)
	}
	return nil
}

func (r *repository) UpdateInvitationStatus(
	ctx context.Context,
	id string,
	status string,
) error {
	_, err := r.db.Exec(
		ctx,
		`
			update invitations
			set status = $2
			where id = $1
		`,
		id,
		status,
	)

	if err != nil {
		return fmt.Errorf("update invitation status: %w", err)
	}
	return nil
}

func (r *repository)RejectInvitation(
	ctx context.Context,
	id string,
) error {

	_, err := r.db.Exec(
		ctx,
		`
			update invitations
			set status = 'rejected'
			where id = $1
			and status = 'pending'	
		`,
		id,
	)
	if err != nil {
		return fmt.Errorf("reject invitation: %w", err)
	}
	return nil
}