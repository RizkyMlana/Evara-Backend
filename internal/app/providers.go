package app

import (
	"evara-backend/internal/family"
	"evara-backend/internal/transaction"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Providers struct {

	TransactionHandler *transaction.Handler

	FamilyHandler *family.Handler

	// UserHandler *user.Handler
	// DashboardHandler *dashboard.Handler
}

func NewProviders(
	db *pgxpool.Pool,
) *Providers {

	// Transaction

	transactionRepo := transaction.NewRepository(db)
	transactionService := transaction.NewService(
		transactionRepo,
	)
	transactionHandler := transaction.NewHandler(
		transactionService,
	)
	// Family

	familyRepo := family.NewRepository(db)

	familyService := family.NewService(
		familyRepo,
	)

	familyHandler := family.NewHandler(
		familyService,
	)


	return &Providers{
		TransactionHandler: transactionHandler,
		FamilyHandler: familyHandler,
	}
}