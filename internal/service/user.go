package service

import (
	"context"
	v1 "todolist/api/todolist/v1"
	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUsersServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

//用户登录
func (s *UserService) Login(ctx context.Context, req *v1.LoginUserRequest) (reply *v1.LoginUserReply, err error) {
	data, err := s.uc.LoginUser(ctx, req)

	if err != nil {
		return &v1.LoginUserReply{
			Code: 400,
			Msg:  err.Error(),
		}, err
	} else {
		return &v1.LoginUserReply{
			Code: 200,
			Msg:  "ok",
			Data: &v1.User{
				Userid:   data.ID,
				Username: data.Username,
				// Email:     data.Email,
				CreatedAt: data.Create_at,
				// Phone:     data.Phone,
			},
		}, nil
	}
}

//用户注册
func (s *UserService) Register(ctx context.Context, req *v1.RegisterUserRequest) (reply *v1.RegisterUserReply, err error) {

	data, err := s.uc.RegisterUser(ctx, req)

	if err != nil {
		return &v1.RegisterUserReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	} else {
		return &v1.RegisterUserReply{
			Code: 400,
			Msg:  "ok",
			Data: &v1.User{
				Userid:    data.ID,
				Username:  data.Username,
				CreatedAt: data.Create_at,
			},
		}, nil
	}
}
