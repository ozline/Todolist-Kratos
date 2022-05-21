package server

import (
	"context"
	v1 "todolist/api/todolist/v1"
	"todolist/internal/conf"
	"todolist/internal/pkg/middleware/auth"
	"todolist/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewSkipRoutersMatcher() selector.MatchFunc {

	skipRouters := map[string]struct{}{
		"/user.v1.Users/Login":    {}, //用户登录
		"/user.v1.Users/Register": {}, //用户注册
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

func NewGRPCServer(c *conf.Server, jwtc *conf.Auth, todolist *service.TodolistService, users *service.UserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			mmd.Server(),
			selector.Server(auth.JWTAuth(jwtc.Secret)).Match(NewSkipRoutersMatcher()).Build(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterTodolistServer(srv, todolist) //注册待办事项服务
	v1.RegisterUsersServer(srv, users)       //注册用户服务
	return srv
}
