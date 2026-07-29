package family

import (
	"context"

	"evara-backend/internal/config"
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
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(
	ctx context.Context,
	userID string,
	name string,
) (string, error) {

	var familyID string

	err := config.DB.QueryRow(
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
		return "", err
	}

	return familyID, nil
}

func (r *repository) AddOwner(
	ctx context.Context,
	familyID string,
	userID string,
) error {

	_, err := config.DB.Exec(
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

	return err
}

func (r *repository) GetMyFamily(
	ctx context.Context,
	userID string,
) (*GetMyFamilyResponse, error) {
	var family GetMyFamilyResponse

	err := config.DB.QueryRow(
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

	if err != nil {
		return nil, ErrFamilyNotFound
	}

	return &family, nil
}

func (r *repository) GetMembers(
	ctx context.Context,
	familyID string,
)([]FamilyMemberResponse, error){
	rows, err := config.DB.Query(
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
		return nil, err
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
			return nil, err
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil

}
