package srv

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_grpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_http "github.com/go-kratos/kratos/v2/transport/http"
)

// new grpc server

func NewGRPCServer(_ log.Logger, network string, addr string, timeout time.Duration, rs func(g *_grpc.Server)) *_grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			//logging.Server(logger),
			// tracing.Server(),
		),
	}
	if network != "" {
		opts = append(opts, grpc.Network(network))
	}
	if addr != "" {
		opts = append(opts, grpc.Address(addr))
	}
	if timeout > 0 {
		opts = append(opts, grpc.Timeout(timeout))
	}
	srv := grpc.NewServer(opts...)
	rs(srv)
	return srv
}

// new http server

func NewHTTPServer(_ log.Logger, network string, addr string, timeout time.Duration, rs func(g *_http.Server)) *_http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			// logging.Server(logger),
		),
	}
	if network != "" {
		opts = append(opts, http.Network(network))
	}
	if addr != "" {
		opts = append(opts, http.Address(addr))
	}
	if timeout > 0 {
		opts = append(opts, http.Timeout(timeout))
	}
	srv := http.NewServer(opts...)
	rs(srv)
	return srv
}
