package biz

import (
	"context"
	v1 "todolist/api/todolist/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type TodolistRepo interface {
	AddTodo(ctx context.Context, g *v1.AddTodoRequest) error
	UpdateTodo(ctx context.Context, g *v1.UpdateTodoRequest) error
	DeleteTodo(ctx context.Context, g *v1.DeleteTodoRequest) error
}

// TodolistUsecase is a Todolist usecase.
type TodolistUsecase struct {
	tr  TodolistRepo
	log *log.Helper
}

type Todo struct {
	UserID    int64
	status    int64
	Title     string
	Message   string
	Create_at int64
	Update_at int64
	Dleete_at int64
	End_at    int64
}

// NewTodolistUsecase new a Todolist usecase.
func NewTodolistUsecase(repo TodolistRepo, logger log.Logger) *TodolistUsecase {
	return &TodolistUsecase{tr: repo, log: log.NewHelper(logger)}
}

func (tc *TodolistUsecase) AddTodo(ctx context.Context, req *v1.AddTodoRequest) error {
	if len(req.Title) == 0 || len(req.Message) == 0 || len(req.Endtime) == 0 {
		return errors.New(422, "Params", "Missing Params Data")
	}

	err := tc.tr.AddTodo(ctx, req)

	return err
}

func (tc *TodolistUsecase) UpdateTodo(ctx context.Context, req *v1.UpdateTodoRequest) error {
	err := tc.tr.UpdateTodo(ctx, req)
	return err
}

func (tc *TodolistUsecase) DeleteTodo(ctx context.Context, req *v1.DeleteTodoRequest) error {
	err := tc.tr.DeleteTodo(ctx, req)
	return err
}

func (tc *TodolistUsecase) ShowAllTodo(ctx context.Context) error {
	return nil
}

func (tc *TodolistUsecase) ShowKeyTodo(ctx context.Context) error {
	return nil
}
