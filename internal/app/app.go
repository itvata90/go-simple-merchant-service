package app

import (
	"context"

	v "github.com/core-go/core/v10"
	"github.com/core-go/health"
	"github.com/core-go/log"
	q "github.com/core-go/sql"
	_ "github.com/go-sql-driver/mysql"

	merchantHandler "merchant-service/internal/merchant/adapter/handler"
	merchantRepository "merchant-service/internal/merchant/adapter/repository"
	merchantPort "merchant-service/internal/merchant/port"
	merchantService "merchant-service/internal/merchant/service"

	teamMemberHandler "merchant-service/internal/team_member/adapter/handler"
	teamMemberRepository "merchant-service/internal/team_member/adapter/repository"
	teamMemberPort "merchant-service/internal/team_member/port"
	teamMemberService "merchant-service/internal/team_member/service"
)

type ApplicationContext struct {
	Health     *health.Handler
	Merchant   merchantPort.MerchantHandler
	TeamMember teamMemberPort.TeamMemberHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	db, err := q.OpenByConfig(conf.Sql)
	if err != nil {
		return nil, err
	}
	logError := log.LogError
	validator := v.NewValidator()

	// merchants
	merchantRepository := merchantRepository.NewMerchantSQLAdapter(db)
	merchantService := merchantService.NewMerchantService(merchantRepository)
	merchantHandler := merchantHandler.NewMerchantHandler(merchantService, validator.Validate, logError)

	// members
	teamMemberRepository := teamMemberRepository.NewTeamMemberAdapter(db)
	teamMemberService := teamMemberService.NewTeamMemberService(teamMemberRepository)
	teamMemberHandler := teamMemberHandler.NewTeamMemberHandler(teamMemberService, validator.Validate, logError)

	sqlChecker := q.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		Health:     healthHandler,
		Merchant:   merchantHandler,
		TeamMember: teamMemberHandler,
	}, nil
}
