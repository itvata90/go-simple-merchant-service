package port

import "net/http"

type MerchantHandler interface {
	GetMerchants(w http.ResponseWriter, r *http.Request)
	GetMerchantByCode(w http.ResponseWriter, r *http.Request)
	CreateMerchant(w http.ResponseWriter, r *http.Request)
	UpdateMerchant(w http.ResponseWriter, r *http.Request)
	DeleteMerchant(w http.ResponseWriter, r *http.Request)
}
