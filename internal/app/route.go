package app

import (
	"context"

	. "github.com/core-go/core"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}
	r.HandleFunc("/health", app.Health.Check).Methods(GET)

	merchantRouter := r.PathPrefix("/merchants").Subrouter()
	merchantRouter.HandleFunc("", app.Merchant.GetMerchants).Methods(GET)
	merchantRouter.HandleFunc("/{code}", app.Merchant.GetMerchantByCode).Methods(GET)
	merchantRouter.HandleFunc("", app.Merchant.CreateMerchant).Methods(POST)
	merchantRouter.HandleFunc("/{code}", app.Merchant.UpdateMerchant).Methods(PUT)
	merchantRouter.HandleFunc("/{code}", app.Merchant.DeleteMerchant).Methods(DELETE)

	memberRouter := r.PathPrefix("/team_members").Subrouter()
	memberRouter.HandleFunc("", app.TeamMember.GetTeamMembers).Methods(GET)
	memberRouter.HandleFunc("/{id}", app.TeamMember.GetTeamMemberById).Methods(GET)
	memberRouter.HandleFunc("/merchants/{merchantCode}", app.TeamMember.GetTeamMemberByMerchantCode).Methods(GET)
	memberRouter.HandleFunc("", app.TeamMember.CreateTeamMember).Methods(POST)
	memberRouter.HandleFunc("/{id}", app.TeamMember.UpdateTeamMember).Methods(PUT)
	memberRouter.HandleFunc("/{id}", app.TeamMember.DeleteTeamMember).Methods(DELETE)

	return nil
}
