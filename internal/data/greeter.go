package data

import (
	"context"

	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewTodolistRepo .
func NewTodolistRepo(data *Data, logger log.Logger) biz.TodolistRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Todolist) (*biz.Todolist, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Todolist) (*biz.Todolist, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Todolist, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Todolist, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Todolist, error) {
	return nil, nil
}
