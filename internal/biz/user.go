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
	Username string
}

type UserRepo interface {
	LoginUser(ctx context.Context, user *User) error
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) LoginUser(ctx context.Context, u *User) error {
	if err := uc.ur.LoginUser(ctx, u); err != nil {

	}
	return nil
}
