package transaction

import (
	"context"
	"errors"
	"evara-backend/internal/config"

	"github.com/jackc/pgx/v5"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(
	ctx context.Context,
	tx Transaction,
) (string, error) {

	var id string

	err := config.DB.QueryRow(
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
		return "", err
	}

	return id, nil
}

func (r *Repository) isFamilyMember(
	ctx context.Context,
	familyID string,
	userID string,
) (bool,error) {
	var exists bool
	err := config.DB.QueryRow(
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
		return  false, err
	}

	return  exists, nil
}

func (r *Repository) GetTransactions(
	ctx context.Context,
	familyID string,
) ([]GetTransactionsResponse, error) {
	rows, err := config.DB.Query(
		ctx,
		`
			select
				t.id,
				COALESCE(p.name, '')
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
		return  nil, err
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
			return  nil, err
		}

		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func (r *Repository) GetByID(
	ctx context.Context,
	id string,
) (*Transaction, error) {
	var tx Transaction

	err := config.DB.QueryRow(
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
	}

	return  &tx, nil
}

func (r *Repository) Delete(
	ctx context.Context,
	id string,
) error {
	_, err := config.DB.Exec(
		ctx,
		`
			delete from transactions
			where id = $1	
		`,
		id,
	)
	return err
}