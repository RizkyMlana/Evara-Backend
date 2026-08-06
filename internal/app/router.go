package app

import (
	"evara-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	p *Providers,
) *gin.Engine {
	r := gin.Default()

	r.SetTrustedProxies(
		[]string{"192.168.1.2"},
	)
	r.Use(
		middleware.RequestID(),
	)
	r.Use(
		middleware.Logging(),
	)
	r.Use(
		middleware.Recovery(),
	)
	api := r.Group("/api")
	
	api.Use(
		middleware.AuthMiddleware(),
	)


	api.GET(
		"/me",
		p.UserHandler.Me,
	)
	api.POST(
		"/families",
		p.FamilyHandler.CreateFamily,
	)

	api.GET(
		"/families/me",
		p.FamilyHandler.GetMyFamily,
	)
	api.GET(
		"/families/:id/members",
		p.FamilyHandler.GetFamilyMember,
	)
	api.POST(
		"/families/:id/invite",
		p.FamilyHandler.InviteMember,
	)

	api.GET(
		"/invitations",
		p.FamilyHandler.GetInvitations,
	)

	api.POST(
		"/invitations/:id/accept",
		p.FamilyHandler.AcceptInvitation,
	)

	api.POST(
		"invitations/:id/reject",
		p.FamilyHandler.RejectInvitation,
	)

	api.POST(
		"/transaction",
		p.TransactionHandler.CreateTransaction,
	)
	api.GET(
		"/transaction",
		p.TransactionHandler.GetTransactions,
	)

	api.DELETE(
		"/transaction/:id",
		p.TransactionHandler.DeleteTransaction,
	)

	api.GET(
		"/dashboard",
		p.DashboardHandler.GetDashboard,
	)
	return r
	
}