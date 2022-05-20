package data

import (
	"context"
	v1 "todolist/api/todolist/v1"
	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewTodolistRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) LoginUser(ctx context.Context, g *v1.LoginUserRequest) (*biz.User, error) {
	// r.data.db.Create()
	return &biz.User{
		Username: "ozline",
		Phone:    15959014518,
		Createat: 13107655114,
		Email:    "ozlinex@outlook.com",
	}, nil
}
