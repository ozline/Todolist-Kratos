package biz

import (
	"context"
	v1 "todolist/api/todolist/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type TodolistRepo interface {
}

// TodolistUsecase is a Todolist usecase.
type TodolistUsecase struct {
	tl  TodolistRepo
	log *log.Helper
}

type Todo struct {
	UserID    int64
	Title     string
	Message   string
	Create_at int64
	Update_at int64
	Dleete_at int64
	End_at    int64
	Token     string
}

// NewTodolistUsecase new a Todolist usecase.
func NewTodolistUsecase(repo TodolistRepo, logger log.Logger) *TodolistUsecase {
	return &TodolistUsecase{tl: repo, log: log.NewHelper(logger)}
}

func (tc *TodolistUsecase) AddTodo(ctx context.Context, req *v1.AddTodoRequest) error {
	if len(req.Title) == 0 || len(req.Message) == 0 || req.Endtime == 0 {
		return errors.New(422, "Params", "Missing Params Data")
	}

	return nil
}

func (tc *TodolistUsecase) DeleteTodo(ctx context.Context) error {
	return nil
}

func (tc *TodolistUsecase) ShowAllTodo(ctx context.Context) error {
	return nil
}

func (tc *TodolistUsecase) ShowKeyTodo(ctx context.Context) error {
	return nil
}

func (tc *TodolistUsecase) ModifyTodo(ctx context.Context) error {
	return nil
}
