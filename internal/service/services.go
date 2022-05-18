package service

import (
	v1 "todolist/api/todolist/v1"
	"todolist/internal/biz"

	"github.com/google/wire"
)

// TodolistService is a greeter service.
type TodolistService struct {
	v1.UnimplementedTodolistServer

	uc *biz.TodolistUsecase
}

// NewTodolistService new a greeter service.
func NewTodolistService(uc *biz.TodolistUsecase) *TodolistService {
	return &TodolistService{uc: uc}
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewTodolistService)
