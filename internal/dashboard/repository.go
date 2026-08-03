package dashboard

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
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
	) (*DashboardResponse, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(
	db *pgxpool.Pool,
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

	return exists, err
}

func (r *repository)GetSummary(
	ctx context.Context,
	familyID string,
) (*DashboardResponse, error) {
	var res DashboardResponse

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
		&res.TotalIncome,
		&res.TotalExpense,
	)

	if err != nil {
		return nil, err
	}
	res.Balance = res.TotalIncome - res.TotalExpense

	return &res, nil
}
