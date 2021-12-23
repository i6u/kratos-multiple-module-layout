package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/i6u/kratos-multiple-module-layout/app/greeter/service/internal/biz"
)

type greeterRepo struct {
	log  *log.Helper
	data *Data
}

func NewGreeterRepo(logger log.Logger, data *Data) biz.GreeterRepo {
	return &greeterRepo{
		log:  log.NewHelper(log.With(logger, "module", "data/greeter")),
		data: data,
	}
}

func (g *greeterRepo) GetGreeter(ctx context.Context, name string) (*biz.Greeter, error) {
	return &biz.Greeter{Name: name}, nil
}
