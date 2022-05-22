package data

import (
	"context"
	v1 "todolist/api/todolist/v1"
	"todolist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	ID         int64
	Username   string
	Password   string
	Status     int64
	Email      string
	Nickname   string
	Phone      int64
	Updated_at int64
	Created_at int64
}

// NewTodolistRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CheckIsUserExist(ctx context.Context, username string) bool {

	var count int64
	_ = r.data.db.Table("users").Where("username = ?", username).Count(&count)
	return count != 0
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (user *biz.User, err error) {
	u := new(User)
	res := r.data.db.Table("users").Where("username = ?", username).First(u)
	if res.Error != nil {
		return nil, err
	}

	return &biz.User{
		Username:  u.Username,
		Status:    u.Status,
		Phone:     u.Phone,
		Create_at: u.Created_at,
		Email:     u.Email,
		Password:  u.Password,
		ID:        u.ID,
	}, nil
}

func (r *userRepo) CreateUser(ctx context.Context, g *v1.RegisterUserRequest) error {

	timeStamp := biz.GetTimestamp13()
	u := User{
		Username:   g.Username,
		Password:   biz.GenerateTokenSHA256(g.Password),
		Phone:      g.Phone,
		Email:      g.Email,
		Created_at: timeStamp,
		Updated_at: 0,
		Nickname:   g.Username,
	}

	res := r.data.db.Table("users").Create(&u)

	return res.Error
}
