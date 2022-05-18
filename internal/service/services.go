package service

import (
	v1 "todolist/api/todolist/v1"
	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewTodolistService)

type TodolistService struct {
	v1.UnimplementedTodolistServer

	uc *biz.UserUsecase
	// tc *biz.TodolistUsecase

	log *log.Helper
}

func NewTodolistService(uc *biz.UserUsecase, logger log.Logger) *TodolistService {
	return &TodolistService{uc: uc, log: log.NewHelper(logger)}
}
