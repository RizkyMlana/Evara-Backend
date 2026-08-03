package app

import (
	"evara-backend/internal/dashboard"
	"evara-backend/internal/family"
	"evara-backend/internal/transaction"
	"evara-backend/internal/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Providers struct {

	TransactionHandler *transaction.Handler

	FamilyHandler *family.Handler

	UserHandler *user.Handler

	DashboardHandler *dashboard.Handler

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

	familyService := family.NewService(familyRepo)

	familyHandler := family.NewHandler(familyService)

	// User

	userService := user.NewService()
	userHandler := user.NewHandler(userService)

	// Dashboard

	dashboardRepo := dashboard.NewRepository(db)
	dashboardService := dashboard.NewService(dashboardRepo)
	dashboardHandler := dashboard.NewHandler(dashboardService)

	return &Providers{
		TransactionHandler: transactionHandler,
		FamilyHandler: familyHandler,
		UserHandler: userHandler,
		DashboardHandler: dashboardHandler,
	}
}