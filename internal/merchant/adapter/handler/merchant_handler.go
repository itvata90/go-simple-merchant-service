package handler

import (
	"context"
	"net/http"
	"reflect"
	"strconv"

	"github.com/google/uuid"

	"github.com/core-go/core"

	domain "merchant-service/internal/merchant/entity"
	"merchant-service/internal/merchant/service"
)

func NewMerchantHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service service.MerchantService, status core.StatusConfig, logError func(context.Context, string, ...map[string]interface{}), validate func(context.Context, interface{}) ([]core.ErrorMessage, error), action *core.ActionConfig) *HttpMerchantHandler {
	modelType := reflect.TypeOf(domain.Merchant{})
	params := core.CreateParams(modelType, &status, logError, validate, action)
	return &HttpMerchantHandler{service: service, Params: params}
}

type HttpMerchantHandler struct {
	service service.MerchantService
	*core.Params
}

func (h *HttpMerchantHandler) GetMerchants(w http.ResponseWriter, r *http.Request) {
	pageIdxParam := r.URL.Query().Get("pageIdx")
	pageSizeParam := r.URL.Query().Get("pageSize")
	pageIdx, err := strconv.Atoi(pageIdxParam)
	if err != nil {
		core.RespondModel(w, r, nil, err, h.Error, nil)
	}
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		core.RespondModel(w, r, nil, err, h.Error, nil)
	}
	res, err := h.service.GetMerchants(r.Context(), pageSize, pageIdx)
	core.RespondModel(w, r, res, err, h.Error, nil)
}

func (h *HttpMerchantHandler) GetMerchantByCode(w http.ResponseWriter, r *http.Request) {
	code := core.GetRequiredParam(w, r)
	if len(code) > 0 {
		res, err := h.service.GetMerchantByCode(r.Context(), code)
		core.RespondModel(w, r, res, err, h.Error, nil)
	}
}

func (h *HttpMerchantHandler) CreateMerchant(w http.ResponseWriter, r *http.Request) {
	var merchant domain.Merchant
	er1 := core.Decode(w, r, &merchant)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &merchant)
		if !core.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			id, err2 := uuid.NewUUID()
			if err2 == nil {
				merchant.Code = id.String()
				res, er3 := h.service.CreateMerchant(r.Context(), &merchant)
				core.AfterCreated(w, r, &merchant, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
			}
		}
	}
}

func (h *HttpMerchantHandler) UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	var merchant domain.Merchant
	er1 := core.DecodeAndCheckId(w, r, &merchant, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &merchant)
		if !core.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, er3 := h.service.UpdateMerchant(r.Context(), &merchant)
			core.HandleResult(w, r, &merchant, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}

func (h *HttpMerchantHandler) DeleteMerchant(w http.ResponseWriter, r *http.Request) {
	code := core.GetRequiredParam(w, r)
	if len(code) > 0 {
		res, err := h.service.DeleteMerchant(r.Context(), code)
		core.HandleDelete(w, r, res, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
