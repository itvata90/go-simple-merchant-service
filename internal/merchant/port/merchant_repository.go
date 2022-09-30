package port

import (
	"context"
	domain "merchant-service/internal/merchant/entity"
)

type MerchantRepository interface {
	GetMerchants(ctx context.Context, pageSize int, pageIdx int) ([]domain.Merchant, int64, error)
	GetMerchantByCode(ctx context.Context, id string) (*domain.Merchant, error)
	CreateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error)
	UpdateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error)
	DeleteMerchant(ctx context.Context, id string) (int64, error)
}
