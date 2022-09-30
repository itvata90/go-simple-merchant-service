package handler

import (
	"context"
	"net/http"
	"reflect"

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

}

func (h *HttpMerchantHandler) GetMerchantByCode(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpMerchantHandler) CreateMerchant(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpMerchantHandler) UpdateMerchant(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpMerchantHandler) DeleteMerchant(w http.ResponseWriter, r *http.Request) {

}
