package service

import (
	v1 "todolist/api/todolist/v1"
	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewTodolistService, NewUserService)

type TodolistService struct {
	v1.UnimplementedTodolistServer
	tc *biz.TodolistUsecase

	log *log.Helper
}

type UserService struct {
	v1.UnimplementedUsersServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func NewTodolistService(tc *biz.TodolistUsecase, logger log.Logger) *TodolistService {
	return &TodolistService{tc: tc, log: log.NewHelper(logger)}
}
