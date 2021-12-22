package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	apiV1 "github.com/i6u/kratos-multiple-module-layout/api/greeter/service/v1"
	"github.com/i6u/kratos-multiple-module-layout/app/greeter/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewGreeterService)

type GreeterService struct {
	apiV1.UnimplementedGreeterServer
	log *log.Helper
	guc *biz.GreeterUseCase
}

func NewGreeterService(logger log.Logger, guc *biz.GreeterUseCase) *GreeterService {
	return &GreeterService{
		log: log.NewHelper(log.With(logger, "moudle", "service/greeter")),
		guc: guc,
	}
}
