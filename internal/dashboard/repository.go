package dashboard

import (
	"context"
	"evara-backend/internal/database"
	"fmt"
)

type Repository interface {
	IsFamilyMember (
		ctx context.Context,
		familyID string,
		userID string,
	) (bool, error)

	GetSummary(
		ctx context.Context,
		familyID string,
	) (*Dashboard, error)
}

type repository struct {
	db database.DBTX
}

func NewRepository(
	db database.DBTX,
) Repository { 
	return &repository{
		db: db,
	}
}

func (r *repository) IsFamilyMember(
	ctx context.Context,
	familyID string,
	userID string,
) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		ctx,
		`
			select exists(
			select 1
			from family_members
			where family_id = $1
			and user_id = $2
			)
		`,
		familyID,
		userID,
	).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("check family member: %w", err)
	}

	return exists, nil
}

func (r *repository)GetSummary(
	ctx context.Context,
	familyID string,
) (*Dashboard, error) {
	var dashboard Dashboard

	err := r.db.QueryRow(
		ctx,
		`
			select 
			coalesce(
				sum(case when type='income' then amount end),0
			),
			coalesce(
				sum(case when type='expense' then amount end),0
			)
			from transactions
			where family_id = $1
		`,
		familyID,
	).Scan(
		&dashboard.TotalIncome,
		&dashboard.TotalExpense,
	)

	if err != nil {
		return nil, fmt.Errorf("get dashboard summary: %w", err)
	}

	return &dashboard, nil
}
