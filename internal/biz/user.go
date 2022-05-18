package biz

import (
	v1 "todolist/api/todolist/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type UserRepo interface {
}

type UserUsecase struct {
	tl  TodolistRepo
	log *log.Helper
}

func NewUserUsecase(repo TodolistRepo, logger log.Logger) *TodolistUsecase {
	return &TodolistUsecase{tl: repo, log: log.NewHelper(logger)}
}

// func (uc *TodolistUsecase) AddTodo(ctx context.Context) error {
// 	return nil
// }
