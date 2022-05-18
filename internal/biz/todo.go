package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type TodolistRepo interface {
}

// TodolistUsecase is a Todolist usecase.
type TodolistUsecase struct {
	tl  TodolistRepo
	log *log.Helper
}

// NewTodolistUsecase new a Todolist usecase.
func NewTodolistUsecase(repo TodolistRepo, logger log.Logger) *TodolistUsecase {
	return &TodolistUsecase{tl: repo, log: log.NewHelper(logger)}
}

func (uc *TodolistUsecase) AddTodo(ctx context.Context) error {
	return nil
}
