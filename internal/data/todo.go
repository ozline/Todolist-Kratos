package data

import (
	"context"
	"strconv"
	v1 "todolist/api/todolist/v1"
	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type todolistRepo struct {
	data *Data
	log  *log.Helper
}

type Todo struct {
	Userid   int64
	Status   int64
	Title    string
	Message  string
	CreateAt int64
	UpdateAt int64
	End_at   int64
}

func getUserid(ctx context.Context) int64 {
	userid, err := strconv.ParseInt(ctx.Value("id").(string), 10, 64)
	if err != nil {
		panic(err)
	}
	return userid
}

func NewTodolistRepo(data *Data, logger log.Logger) biz.TodolistRepo {
	return &todolistRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *todolistRepo) AddTodo(ctx context.Context, g *v1.AddTodoRequest) error {
	endat, _ := strconv.ParseInt(g.Endtime, 10, 64)
	t := Todo{
		Status:   0,
		Userid:   getUserid(ctx),
		Title:    g.Title,
		Message:  g.Message,
		CreateAt: biz.GetTimestamp13(),
		End_at:   endat,
	}

	res := r.data.db.Table("todolist").Create(&t)
	return res.Error
}

func (r *todolistRepo) UpdateTodo(ctx context.Context, g *v1.UpdateTodoRequest) error {

	res := r.data.db.Table("todolist").Where("userid = ? AND id = ?", getUserid(ctx), g.Id).Update("status", g.Status)
	return res.Error
}

func (r *todolistRepo) DeleteTodo(ctx context.Context, g *v1.DeleteTodoRequest) error {
	u := new(Todo)
	res := r.data.db.Table("todolist").Where("userid = ? AND id = ?", getUserid(ctx), g.Id).Unscoped().Delete(&u)
	return res.Error
}
