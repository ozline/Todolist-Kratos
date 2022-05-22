package service

import (
	"context"
	v1 "todolist/api/todolist/v1"
	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type TodolistService struct {
	v1.UnimplementedTodolistServer
	tc *biz.TodolistUsecase

	log *log.Helper
}

func NewTodolistService(tc *biz.TodolistUsecase, logger log.Logger) *TodolistService {
	return &TodolistService{tc: tc, log: log.NewHelper(logger)}
}

//添加待办事项
func (s *TodolistService) Add(ctx context.Context, req *v1.AddTodoRequest) (reply *v1.AddTodoReply, err error) {

	err = s.tc.AddTodo(ctx, req)
	if err != nil {
		return &v1.AddTodoReply{
			Code: 400,
			Msg:  err.Error(),
		}, err
	}
	return &v1.AddTodoReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

//编辑待办事项
func (s *TodolistService) Update(ctx context.Context, req *v1.UpdateTodoRequest) (reply *v1.UpdateTodoReply, err error) {

	err = s.tc.UpdateTodo(ctx, req)

	if err != nil {
		return &v1.UpdateTodoReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &v1.UpdateTodoReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

//删除待办事项
func (s *TodolistService) Delete(ctx context.Context, req *v1.DeleteTodoRequest) (reply *v1.DeleteTodoReply, err error) {
	err = s.tc.DeleteTodo(ctx, req)

	if err != nil {
		return &v1.DeleteTodoReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &v1.DeleteTodoReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

//显示全部待办事项
func (s *TodolistService) ShowAll(ctx context.Context, req *v1.ShowAllTodoRequest) (reply *v1.ShowAllTodoReply, err error) {
	return &v1.ShowAllTodoReply{}, nil
}

//根据关键词查找待办事项
func (s *TodolistService) ShowKey(ctx context.Context, req *v1.ShowKeyTodoRequest) (reply *v1.ShowKeyTodoReply, err error) {
	return &v1.ShowKeyTodoReply{}, nil
}
