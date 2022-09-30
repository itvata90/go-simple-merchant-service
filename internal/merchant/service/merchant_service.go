package service

import (
	"context"
	domain "merchant-service/internal/merchant/entity"
	"merchant-service/internal/merchant/port"
)

type MerchantService interface {
	GetMerchants(ctx context.Context, pageSize int, pageIdx int) ([]domain.Merchant, error)
	GetMerchantByCode(ctx context.Context, code string) (*domain.Merchant, error)
	CreateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error)
	UpdateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error)
	DeleteMerchant(ctx context.Context, code string) (int64, error)
}

func NewMerchantService(repository port.MerchantRepository) MerchantService {
	return &merchantService{repository: repository}
}

type merchantService struct {
	repository port.MerchantRepository
}

func (s *merchantService) GetMerchants(ctx context.Context, pageSize int, pageIdx int) ([]domain.Merchant, error) {
	return s.repository.GetMerchants(ctx, pageSize, pageIdx)
}

func (s *merchantService) GetMerchantByCode(ctx context.Context, code string) (*domain.Merchant, error) {
	return s.repository.GetMerchantByCode(ctx, code)
}

func (s *merchantService) CreateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error) {
	return s.repository.CreateMerchant(ctx, merchant)
}

func (s *merchantService) UpdateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error) {
	return s.repository.UpdateMerchant(ctx, merchant)
}

func (s *merchantService) DeleteMerchant(ctx context.Context, code string) (int64, error) {
	return s.repository.DeleteMerchant(ctx, code)
}
