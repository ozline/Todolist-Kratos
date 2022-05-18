package data

import (
	"context"
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

func (r *userRepo) LoginUser(ctx context.Context, g *biz.User) error {
	// r.data.db.Create()
	return nil
}
