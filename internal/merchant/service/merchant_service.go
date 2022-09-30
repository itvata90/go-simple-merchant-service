package service

import (
	"context"
	domain "merchant-service/internal/merchant/entity"
)

type MerchantService interface {
	GetMerchants(ctx context.Context, pageSize int, pageIdx int) ([]domain.Merchant, error)
	GetMerchantByCode(ctx context.Context, code string) (*domain.Merchant, error)
	CreateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error)
	UpdateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error)
	DeleteMerchant(ctx context.Context, code string) (int64, error)
}
