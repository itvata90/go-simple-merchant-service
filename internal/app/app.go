package app

import (
	"context"
	"reflect"

	sv "github.com/core-go/core"
	v "github.com/core-go/core/v10"
	"github.com/core-go/health"
	"github.com/core-go/log"
	"github.com/core-go/search/query"
	q "github.com/core-go/sql"
	_ "github.com/go-sql-driver/mysql"

	merchantHandler "merchant-service/internal/merchant/adapter/handler"
	merchantRepository "merchant-service/internal/merchant/adapter/repository"
	merchantDomain "merchant-service/internal/merchant/entity"
	merchantPort "merchant-service/internal/merchant/port"
	merchantService "merchant-service/internal/merchant/service"
)

type ApplicationContext struct {
	Health   *health.Handler
	Merchant merchantPort.MerchantHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	db, err := q.OpenByConfig(conf.Sql)
	if err != nil {
		return nil, err
	}
	logError := log.LogError
	status := sv.InitializeStatus(conf.Status)
	action := sv.InitializeAction(conf.Action)
	validator := v.NewValidator()

	// merchants
	merchantType := reflect.TypeOf(merchantDomain.Merchant{})
	merchantQueryBuilder := query.NewBuilder(db, "merchants", merchantType)
	merchantSearchBuilder, err := q.NewSearchBuilder(db, merchantType, merchantQueryBuilder.BuildQuery)
	if err != nil {
		return nil, err
	}
	merchantRepository := merchantRepository.NewMerchantSQLAdapter(db)
	merchantService := merchantService.NewMerchantService(merchantRepository)
	merchantHandler := merchantHandler.NewMerchantHandler(merchantSearchBuilder.Search, merchantService, status, logError, validator.Validate, &action)

	// members

	sqlChecker := q.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		Health:   healthHandler,
		Merchant: merchantHandler,
	}, nil
}
