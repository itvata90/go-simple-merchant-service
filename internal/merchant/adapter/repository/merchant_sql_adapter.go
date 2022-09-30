package repository

import (
	"context"
	"database/sql"
	domain "merchant-service/internal/merchant/entity"
)

func NewMerchantSQLAdapter(db *sql.DB) *MerchantSQLAdapter {
	return &MerchantSQLAdapter{DB: db}
}

type MerchantSQLAdapter struct {
	DB *sql.DB
}

func (r *MerchantSQLAdapter) GetMerchants(ctx context.Context, pageSize int, pageIdx int) ([]domain.Merchant, error) {
	return nil, nil
}

func (r *MerchantSQLAdapter) GetMerchantByCode(ctx context.Context, code string) (*domain.Merchant, error) {
	return nil, nil
}

func (r *MerchantSQLAdapter) CreateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error) {
	return 0, nil
}

func (r *MerchantSQLAdapter) UpdateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error) {
	return 0, nil
}

func (r *MerchantSQLAdapter) DeleteMerchant(ctx context.Context, code string) (int64, error) {
	return 0, nil
}
