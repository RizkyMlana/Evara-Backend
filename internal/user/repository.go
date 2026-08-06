package user

import (
	"context"
	"errors"
	"evara-backend/internal/database"
	"evara-backend/pkg/apperror"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	GetByID(
		ctx context.Context,
		id string,
	) (*User, error) 
}

type repository struct {
	db database.DBTX
}

func NewRepository(
	db database.DBTX,
) Repository {
	return  &repository{
		db: db,
	}
}

func (r *repository) GetByID(
	ctx context.Context,
	id string,
) (*User, error) {
	var user User

	err := r.db.QueryRow(
		ctx,
		`select id, name, email
		from profiles
		where id = $1
		`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, apperror.NotFound(
			"user not found",
		)
	}

	if err != nil {
		return nil, fmt.Errorf("get profiles: %w", err)
	}

	return &user, nil
}