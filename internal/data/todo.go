package data

import (
	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type todolistRepo struct {
	data *Data
	log  *log.Helper
}

// NewTodolistRepo .
func NewTodolistRepo(data *Data, logger log.Logger) biz.TodolistRepo {
	return &todolistRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// func (r *greeterRepo) Save(ctx context.Context, g *biz.Todolist) (*biz.Todolist, error) {
// 	return g, nil
// }
