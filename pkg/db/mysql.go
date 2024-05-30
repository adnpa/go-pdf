package db

import (
	"context"
	"fmt"
	conf "github.com/adnpa/gpdf/config"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init 初始化Postgres连接
func init() {
	var err error

	cfg := conf.Cfg.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AllowGlobalUpdate = true
	sqldb, _ := db.DB()
	sqldb.SetMaxOpenConns(conf.Cfg.MaxOpenConns)
	sqldb.SetMaxIdleConns(conf.Cfg.MaxIdleConns)
	migration()
	return
}

// Close 关闭数据库连接
func Close() {
	sqldb, _ := db.DB()
	sqldb.Close()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
