package port

import (
	"context"
	domain "merchant-service/internal/team_member/entity"
)

type TeamMemberRepository interface {
	GetTeamMembers(ctx context.Context, pageSize int, pageIndex int) ([]domain.TeamMember, int64, error)
	GetTeamMemberById(ctx context.Context, id string) (*domain.TeamMember, error)
	GetTeamMemberByMerchantCode(ctx context.Context, merchantCode string, pageSize int, pageIndex int) ([]domain.TeamMember, int64, error)
	CreateTeamMember(ctx context.Context, merchant *domain.TeamMember) (int64, error)
	UpdateTeamMember(ctx context.Context, merchant *domain.TeamMember) (int64, error)
	DeleteTeamMember(ctx context.Context, id string) (int64, error)
}
