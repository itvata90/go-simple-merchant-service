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

	merchant := "/merchants"
	r.HandleFunc(merchant, app.Merchant.GetMerchants).Methods(GET)
	r.HandleFunc(merchant+"/{code}", app.Merchant.GetMerchantByCode).Methods(GET)
	r.HandleFunc(merchant, app.Merchant.CreateMerchant).Methods(POST)
	r.HandleFunc(merchant+"/{code}", app.Merchant.UpdateMerchant).Methods(PUT)
	r.HandleFunc(merchant+"/{code}", app.Merchant.DeleteMerchant).Methods(DELETE)

	return nil
}
