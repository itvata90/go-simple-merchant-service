package service

import (
	"context"
	domain "merchant-service/internal/merchant/entity"
	"merchant-service/internal/merchant/port"
	"time"

	"github.com/google/uuid"
)

type MerchantService interface {
	GetMerchants(ctx context.Context, pageSize int, pageIdx int) (*domain.GetMerchantsResponse, error)
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

func (s *merchantService) GetMerchants(ctx context.Context, pageSize int, pageIndex int) (*domain.GetMerchantsResponse, error) {
	merchants, total, err := s.repository.GetMerchants(ctx, pageSize, pageIndex)
	if err != nil {
		return nil, err
	}
	return &domain.GetMerchantsResponse{
		Data: merchants,
		Paging: domain.Pagination{
			Total:     total,
			PageIndex: pageIndex,
			PageSize:  pageSize,
		},
	}, nil
}

func (s *merchantService) GetMerchantByCode(ctx context.Context, code string) (*domain.Merchant, error) {
	return s.repository.GetMerchantByCode(ctx, code)
}

func (s *merchantService) CreateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error) {
	t := time.Now()
	merchant.CreatedAt = &t
	merchant.Code = uuid.New().String()
	return s.repository.CreateMerchant(ctx, merchant)
}

func (s *merchantService) UpdateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error) {
	t := time.Now()
	merchant.UpdatedAt = &t
	return s.repository.UpdateMerchant(ctx, merchant)
}

func (s *merchantService) DeleteMerchant(ctx context.Context, code string) (int64, error) {
	return s.repository.DeleteMerchant(ctx, code)
}
