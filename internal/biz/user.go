package biz

import (
	"context"
	v1 "todolist/api/todolist/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type User struct {
	Username  string
	Email     string
	Phone     int64
	Create_at int64
	Password  string
	Nickname  string
	Token     string
}

type UserRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, req *v1.RegisterUserRequest) error
	CheckIsUserExist(ctx context.Context, username string) bool
	// GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) LoginUser(ctx context.Context, u *v1.LoginUserRequest) (*User, error) {
	if len(u.Username) == 0 || len(u.Password) == 0 {
		return nil, errors.New(422, "Params", "Missing Params Data")
	}
	data, err := uc.ur.GetUserByUsername(ctx, u.Username)
	if err != nil {
		return nil, err
	}
	if GenerateTokenSHA256(u.Password) != data.Password {
		return nil, errors.New(422, "Params", "username of password invalid")
	}

	return data, err
}

func (uc *UserUsecase) RegisterUser(ctx context.Context, u *v1.RegisterUserRequest) (*User, error) {
	if len(u.Username) == 0 || len(u.Password) == 0 || u.Phone == 0 || len(u.Email) == 0 {
		return nil, errors.New(422, "Params", "Missing Params Data")
	}

	//检查用户名是否重复
	if uc.ur.CheckIsUserExist(ctx, u.Username) {
		return nil, errors.New(422, "Params", "User Existed")
	}

	//创建用户
	err := uc.ur.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	//注册成功后调用一次获取用户信息的API，回调用户信息
	data, err := uc.ur.GetUserByUsername(ctx, u.Username)

	if err != nil {
		return nil, err
	}

	return data, nil
}
