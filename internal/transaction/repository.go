package transaction

import (
	"context"
	"errors"
	"evara-backend/internal/database"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)


type Repository interface {
	Create(
		ctx context.Context,
		tx Transaction,
	) (string, error)

	IsFamilyMember(
		ctx context.Context,
		familyID string,
		userID string,
	) (bool, error)

	GetTransactions(
		ctx context.Context,
		familyID string,
	) ([]GetTransactionsResponse, error)

	GetByID(
		ctx context.Context,
		id string,
	) (*Transaction, error)

	Delete(
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
	tx Transaction,
) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var id string

	err := r.db.QueryRow(
		ctx,
		`
		insert into transactions (
			family_id,
			user_id,
			type,
			title,
			description,
			amount
		)
		values ($1,$2,$3,$4,$5,$6)
		returning id
		`,
		tx.FamilyID,
		tx.UserID,
		tx.Type,
		tx.Title,
		tx.Description,
		tx.Amount,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("create transaction: %w", err)
	}

	return id, nil
}

func (r *repository) IsFamilyMember(
	ctx context.Context,
	familyID string,
	userID string,
) (bool,error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var exists bool
	err := r.db.QueryRow(
		ctx,
		`select exists (
			select 1
			from family_members
			where family_id = $1
			and user_id = $2
		)`,
		familyID,
		userID,
	).Scan(&exists)

	if err != nil {
		return  false, fmt.Errorf("check family member: %w", err)
	}

	return  exists, nil
}

func (r *repository) GetTransactions(
	ctx context.Context,
	familyID string,
) ([]GetTransactionsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	rows, err := r.db.Query(
		ctx,
		`
			select
				t.id,
				COALESCE(p.name, ''),
				t.type,
				t.title,
				t.description,
				t.amount,
				t.created_at
			from transactions t
			join profiles p on p.id = t.user_id
			where	t.family_id = $1
			order by t.created_at desc
		`,
		familyID,
	)

	if err != nil {
		return  nil, fmt.Errorf("query transactions: %w", err)
	}

	defer rows.Close()
	
	var transactions []GetTransactionsResponse

	for rows.Next() {
		var tx GetTransactionsResponse

		err := rows.Scan(
			&tx.ID,
			&tx.UserName,
			&tx.Type,
			&tx.Title,
			&tx.Description,
			&tx.Amount,
			&tx.CreatedAt,
		)

		if err != nil {
			return  nil, fmt.Errorf("scan transaction: %w", err)
		}

		transactions = append(transactions, tx)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate transactions: %w", err)
	}

	return transactions, nil
}

func (r *repository) GetByID(
	ctx context.Context,
	id string,
) (*Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var tx Transaction

	err := r.db.QueryRow(
		ctx,
		`
		select id, family_id, user_id, type, title, description, amount, created_at
		from transactions
		where id = $1
		`,
		id,
	).Scan(
		&tx.ID,
		&tx.FamilyID,
		&tx.UserID,
		&tx.Type,
		&tx.Title,
		&tx.Description,
		&tx.Amount,
		&tx.CreatedAt,
	)
	if err != nil{
		if errors.Is(err, pgx.ErrNoRows) {
			return  nil, ErrTransactionNotFound
		}
		return nil, fmt.Errorf("get transaction by id: %w", err)
	}

	return  &tx, nil
}

func (r *repository) Delete(
	ctx context.Context,
	id string,
) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := r.db.Exec(
		ctx,
		`
			delete from transactions
			where id = $1	
		`,
		id,
	)
	if err != nil {
		return fmt.Errorf("delete transaction: %w", err)
	}

	return nil
}