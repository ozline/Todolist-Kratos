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

type User struct {
	Username   string `gorm:"size:500"`
	Password   string `gorm:"size:500"`
	Status     int64  `gorm:"size:500"`
	Email      string `gorm:"size:500"`
	Nickname   string `gorm:"size:500"`
	Phone      int64  `gorm:"size:500"`
	Updated_at int64  `gorm:"size:500"`
	Created_at int64  `gorm:"size:500"`
	// Deleted_at int64  `gorm:"size:500,default:''"`
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
		Phone:     u.Phone,
		Create_at: u.Created_at,
		Email:     u.Email,
		Password:  u.Password,
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
