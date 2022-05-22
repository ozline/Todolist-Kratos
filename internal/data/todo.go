package data

import (
	"context"
	"strconv"
	v1 "todolist/api/todolist/v1"
	"todolist/internal/biz"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type todolistRepo struct {
	data *Data
	log  *log.Helper
}

type Todo struct {
	Userid   int64
	Status   int64
	Title    string
	Message  string
	CreateAt int64
	UpdateAt int64
	End_at   int64
}

func getUserid(ctx context.Context) int64 {
	userid, err := strconv.ParseInt(ctx.Value("id").(string), 10, 64)
	if err != nil {
		panic(err)
	}
	return userid
}

func NewTodolistRepo(data *Data, logger log.Logger) biz.TodolistRepo {
	return &todolistRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *todolistRepo) AddTodo(ctx context.Context, g *v1.AddTodoRequest) error {
	endat, _ := strconv.ParseInt(g.Endtime, 10, 64)
	t := Todo{
		Status:   0,
		Userid:   getUserid(ctx),
		Title:    g.Title,
		Message:  g.Message,
		CreateAt: biz.GetTimestamp13(),
		End_at:   endat,
	}

	res := r.data.db.Table("todolist").Create(&t)
	return res.Error
}

func (r *todolistRepo) UpdateTodo(ctx context.Context, g *v1.UpdateTodoRequest) (total int64, err error) {

	var res *gorm.DB

	// id == -1 = 全部已完成设为待办
	// id == -2 = 全部待办设为已完成

	res = r.data.db.Table("todolist")
	if g.Id < 0 {
		res = res.Where("userid = ? AND status = ?", getUserid(ctx), g.Id+2).Update("status", (3+g.Id)%2)
	} else {
		res = res.Where("userid = ? AND id = ?", getUserid(ctx), g.Id).Update("status", g.Status)
	}

	// res = res.Update("update_at", biz.GetTimestamp13())
	spew.Dump(res.RowsAffected)
	return res.RowsAffected, res.Error // 0
}

func (r *todolistRepo) DeleteTodo(ctx context.Context, g *v1.DeleteTodoRequest) (total int64, err error) {
	u := new(Todo)
	var res *gorm.DB

	// id == -3 = 全部
	// id == -2 = 未完成
	// id == -1 = 已完成
	res = r.data.db.Table("todolist")

	if g.Id < 0 {
		if g.Id == -3 {
			res = res.Where("userid = ?", getUserid(ctx))
		} else {
			res = res.Where("userid = ? AND status = ?", getUserid(ctx), g.Id+2)
		}
	} else {
		res = res.Where("userid = ? AND id = ?", getUserid(ctx), g.Id)
	}

	res = res.Unscoped().Delete(&u)

	return res.RowsAffected, res.Error
}

func (r *todolistRepo) GetTodolistByStatus(ctx context.Context, g *v1.ShowAllTodoRequest) (list []*biz.Todo, total int64, err error) {
	var todos []Todo
	var res *gorm.DB
	var count int64
	// status <0 = 获取全部
	res = r.data.db.Table("todolist")

	if g.Status < 0 {
		res = res.Where("userid = ?", getUserid(ctx)).Offset(int((g.Page - 1) * g.Pagesize)).Limit(int(g.Pagesize)).Find(&todos).Count(&count)
	} else {
		res = res.Where("userid = ? AND status = ?", getUserid(ctx), g.Status)
	}

	res = res.Offset(int((g.Page - 1) * g.Pagesize)).Limit(int(g.Pagesize)).Find(&todos).Count(&count)

	if res.Error != nil {
		return nil, 0, res.Error
	}

	list = make([]*biz.Todo, 0)

	for _, x := range todos {
		list = append(list, &biz.Todo{
			Title:     x.Title,
			Message:   x.Message,
			Status:    x.Status,
			End_at:    x.End_at,
			Create_at: x.CreateAt,
			Update_at: x.UpdateAt,
		})
	}
	return list, count, nil
}

func (r *todolistRepo) GetTodolistByKey(ctx context.Context, g *v1.ShowKeyTodoRequest) (list []*biz.Todo, total int64, err error) {
	var todos []Todo
	var count int64
	res := r.data.db.Table("todolist").Where("userid = ? AND message LIKE ?", getUserid(ctx), "%"+g.Key+"%").Offset(int((g.Page - 1) * g.Pagesize)).Limit(int(g.Pagesize)).Find(&todos).Count(&count)

	if res.Error != nil {
		return nil, 0, res.Error
	}

	list = make([]*biz.Todo, 0)

	for _, x := range todos {
		list = append(list, &biz.Todo{
			Title:     x.Title,
			Message:   x.Message,
			Status:    x.Status,
			End_at:    x.End_at,
			Create_at: x.CreateAt,
			Update_at: x.UpdateAt,
		})
	}
	return list, count, nil
}
