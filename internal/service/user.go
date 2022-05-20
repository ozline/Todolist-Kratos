package service

import (
	"context"
	v1 "todolist/api/todolist/v1"
)

//用户登录
func (s *TodolistService) LoginUser(ctx context.Context, req *v1.LoginUserRequest) (reply *v1.LoginUserReply, err error) {
	user, err := s.uc.LoginUser(ctx, req)

	if err != nil && user == nil {
		return &v1.LoginUserReply{
			Code: 400,
			Msg:  err.Error(),
		}, err
	} else {
		return &v1.LoginUserReply{
			Code: 200,
			Msg:  "ok",
			Data: &v1.User{
				Username: user.Username,
				Email:    user.Email,
				Createat: user.Createat,
				Phone:    user.Phone,
			},
		}, nil
	}
}

//用户注册
func (s *TodolistService) RegisterUser(ctx context.Context, req *v1.RegisterUserRequest) (reply *v1.RegisterUserReply, err error) {
	return &v1.RegisterUserReply{}, nil
}
