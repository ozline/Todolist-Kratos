package service

import (
	"context"
	v1 "todolist/api/todolist/v1"
)

//添加待办事项
func (s *TodolistService) AddTodo(ctx context.Context, req *v1.AddTodoRequest) (reply *v1.AddTodoReply, err error) {
	return &v1.AddTodoReply{}, nil
}

//显示全部待办事项
func (s *TodolistService) ShowAllTodo(ctx context.Context, req *v1.ShowAllTodoRequest) (reply *v1.ShowAllTodoReply, err error) {
	return &v1.ShowAllTodoReply{}, nil
}

//根据关键词查找待办事项
func (s *TodolistService) ShowKeyTodo(ctx context.Context, req *v1.ShowKeyTodoRequest) (reply *v1.ShowKeyTodoReply, err error) {
	return &v1.ShowKeyTodoReply{}, nil
}

//删除待办事项
func (s *TodolistService) DeleteTodo(ctx context.Context, req *v1.DeleteTodoRequest) (reply *v1.DeleteTodoReply, err error) {
	return &v1.DeleteTodoReply{}, nil
}

//编辑待办事项
func (s *TodolistService) ModifyTodo(ctx context.Context, req *v1.ModifyTodoRequest) (reply *v1.ModifyTodoReply, err error) {
	return &v1.ModifyTodoReply{}, nil
}
