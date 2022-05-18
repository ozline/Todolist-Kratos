package data

import (
	"strconv"
	"strings"
	"todolist/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewTodolistRepo)

// Data .
type Data struct {
	// TODO wrapped database client

	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	path := strings.Join([]string{c.Database.Username, ":", c.Database.Password, "@tcp(", c.Database.Address, ":", strconv.FormatInt(c.Database.Port, 10), ")/", c.Database.Dbname, "?charset=utf8&parseTime=True"}, "") //数据库连接文本

	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}

	return db
}
