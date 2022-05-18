package service

import (
	"context"
	v1 "todolist/api/todolist/v1"
)

//用户登录
func (s *TodolistService) LoginUser(ctx context.Context, req *v1.LoginUserRequest) (reply *v1.LoginUserReply, err error) {
	return &v1.LoginUserReply{
		Unifiedreply: &v1.UnifiedReply{
			Code: 200,
			Msg:  "ok",
		},
	}, nil
}

//用户注册
func (s *TodolistService) RegisterUser(ctx context.Context, req *v1.RegisterUserRequest) (reply *v1.RegisterUserReply, err error) {
	return &v1.RegisterUserReply{}, nil
}
