package biz

import (
	"context"

	v1 "todolist/api/todolist/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Todolist is a Todolist model.
type Todolist struct {
	Hello string
}

// TodolistRepo is a Greater repo.
type TodolistRepo interface {
	Save(context.Context, *Todolist) (*Todolist, error)
	Update(context.Context, *Todolist) (*Todolist, error)
	FindByID(context.Context, int64) (*Todolist, error)
	ListByHello(context.Context, string) ([]*Todolist, error)
	ListAll(context.Context) ([]*Todolist, error)
}

// TodolistUsecase is a Todolist usecase.
type TodolistUsecase struct {
	repo TodolistRepo
	log  *log.Helper
}

// NewTodolistUsecase new a Todolist usecase.
func NewTodolistUsecase(repo TodolistRepo, logger log.Logger) *TodolistUsecase {
	return &TodolistUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateTodolist creates a Todolist, and returns the new Todolist.
func (uc *TodolistUsecase) CreateTodolist(ctx context.Context, g *Todolist) (*Todolist, error) {
	uc.log.WithContext(ctx).Infof("CreateTodolist: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
