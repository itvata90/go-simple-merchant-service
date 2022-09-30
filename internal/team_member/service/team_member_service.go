package service

import (
	"context"
	domain "merchant-service/internal/team_member/entity"
	"merchant-service/internal/team_member/port"
	"time"

	"github.com/google/uuid"
)

type TeamMemberService interface {
	GetTeamMembers(ctx context.Context, pageSize int, pageIndex int) (*domain.GetTeamMembersResponse, error)
	GetTeamMemberById(ctx context.Context, id string) (*domain.TeamMember, error)
	GetTeamMemberByMerchantCode(ctx context.Context, merchantCode string, pageSize int, pageIndex int) (*domain.GetTeamMembersResponse, error)
	CreateTeamMember(ctx context.Context, teamMember *domain.TeamMember) (int64, error)
	UpdateTeamMember(ctx context.Context, teamMember *domain.TeamMember) (int64, error)
	DeleteTeamMember(ctx context.Context, id string) (int64, error)
}

func NewTeamMemberService(repository port.TeamMemberRepository) TeamMemberService {
	return &teamMemberService{repository: repository}
}

type teamMemberService struct {
	repository port.TeamMemberRepository
}

func (s *teamMemberService) GetTeamMembers(ctx context.Context, pageSize int, pageIndex int) (*domain.GetTeamMembersResponse, error) {
	teamMembers, total, err := s.repository.GetTeamMembers(ctx, pageSize, pageIndex)
	if err != nil {
		return nil, err
	}
	return &domain.GetTeamMembersResponse{
		Data: teamMembers,
		Paging: domain.Pagination{
			Total:     total,
			PageIndex: pageIndex,
			PageSize:  pageSize,
		},
	}, nil
}

func (s *teamMemberService) GetTeamMemberById(ctx context.Context, id string) (*domain.TeamMember, error) {
	return s.repository.GetTeamMemberById(ctx, id)
}

func (s *teamMemberService) GetTeamMemberByMerchantCode(ctx context.Context, merchantCode string, pageSize int, pageIndex int) (*domain.GetTeamMembersResponse, error) {
	teamMembers, total, err := s.repository.GetTeamMemberByMerchantCode(ctx, merchantCode, pageSize, pageIndex)
	if err != nil {
		return nil, err
	}
	return &domain.GetTeamMembersResponse{
		Data: teamMembers,
		Paging: domain.Pagination{
			Total:     total,
			PageIndex: pageIndex,
			PageSize:  pageSize,
		},
	}, nil
}

func (s *teamMemberService) CreateTeamMember(ctx context.Context, teamMember *domain.TeamMember) (int64, error) {
	t := time.Now()
	teamMember.CreatedAt = &t
	teamMember.Id = uuid.New().String()

	return s.repository.CreateTeamMember(ctx, teamMember)
}

func (s *teamMemberService) UpdateTeamMember(ctx context.Context, teamMember *domain.TeamMember) (int64, error) {
	t := time.Now()
	teamMember.UpdatedAt = &t
	return s.repository.UpdateTeamMember(ctx, teamMember)
}

func (s *teamMemberService) DeleteTeamMember(ctx context.Context, id string) (int64, error) {
	return s.repository.DeleteTeamMember(ctx, id)
}
