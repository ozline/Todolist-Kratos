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
	Username     string
	Email        string
	Phone        int64
	Createat     int64
	PasswordHash string
	Nickname     string
	Token        string
}

type UserRepo interface {
	LoginUser(ctx context.Context, req *v1.LoginUserRequest) (*User, error)
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
	data, err := uc.ur.LoginUser(ctx, u)
	if err != nil {
		panic(data)
	}
	return data, nil
}
