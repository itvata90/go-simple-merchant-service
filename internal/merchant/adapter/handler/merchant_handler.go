package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	sv "github.com/core-go/core"
	"github.com/gorilla/mux"

	domain "merchant-service/internal/merchant/entity"
	"merchant-service/internal/merchant/service"
)

func NewMerchantHandler(service service.MerchantService, validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string, ...map[string]interface{})) *HttpMerchantHandler {
	return &HttpMerchantHandler{service: service, validate: validate, logError: logError}
}

type HttpMerchantHandler struct {
	service  service.MerchantService
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error)
	logError func(context.Context, string, ...map[string]interface{})
}

func (h *HttpMerchantHandler) GetMerchants(w http.ResponseWriter, r *http.Request) {
	pageIndexParam := r.URL.Query().Get("pageIndex")
	pageSizeParam := r.URL.Query().Get("pageSize")
	pageIndex, err := strconv.Atoi(pageIndexParam)
	if err != nil {
		http.Error(w, "pageIndex must be an integer", http.StatusBadRequest)
		return
	}
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		http.Error(w, "pageSize must be an integer", http.StatusBadRequest)
		return
	}
	res, err := h.service.GetMerchants(r.Context(), pageSize, pageIndex)
	if err != nil {
		h.logError(r.Context(), err.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *HttpMerchantHandler) GetMerchantByCode(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		http.Error(w, "code cannot be empty", http.StatusBadRequest)
		return
	}
	merchant, err := h.service.GetMerchantByCode(r.Context(), code)
	if err != nil {
		h.logError(r.Context(), err.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	if merchant == nil {
		JSON(w, http.StatusNotFound, nil)
		return
	}
	JSON(w, http.StatusOK, merchant)
}

func (h *HttpMerchantHandler) CreateMerchant(w http.ResponseWriter, r *http.Request) {
	var merchant domain.Merchant
	er1 := json.NewDecoder(r.Body).Decode(&merchant)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	errors, er2 := h.validate(r.Context(), &merchant)
	if er2 != nil {
		h.logError(r.Context(), er2.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		h.logError(r.Context(), er2.Error())
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er3 := h.service.CreateMerchant(r.Context(), &merchant)
	if er3 != nil {
		h.logError(r.Context(), er3.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, res)
}

func (h *HttpMerchantHandler) UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	var merchant domain.Merchant
	er1 := json.NewDecoder(r.Body).Decode(&merchant)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(merchant.Code) == 0 {
		merchant.Code = code
	} else if code != merchant.Code {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	errors, er2 := h.validate(r.Context(), &merchant)
	if er2 != nil {
		h.logError(r.Context(), er2.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		h.logError(r.Context(), er2.Error())
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er3 := h.service.UpdateMerchant(r.Context(), &merchant)
	if er3 != nil {
		h.logError(r.Context(), er3.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *HttpMerchantHandler) DeleteMerchant(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		http.Error(w, "code cannot be empty", http.StatusBadRequest)
		return
	}
	res, err := h.service.DeleteMerchant(r.Context(), code)
	if err != nil {
		h.logError(r.Context(), err.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
