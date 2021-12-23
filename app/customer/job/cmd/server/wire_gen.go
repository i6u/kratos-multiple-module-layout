// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/biz"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/conf"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/data"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/server"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/service"
)

// Injectors from wire.go:

func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	customerRepo := data.NewCustomerRepo(logger)
	customerUseCase := biz.NewCustomerUseCase(logger, customerRepo)
	customerService := service.NewCustomerService(logger, customerUseCase)
	httpServer := server.NewHTTPServer(logger, confServer, customerService)
	grpcServer := server.NewGRPCServer(logger, confServer, customerService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
	}, nil
}
