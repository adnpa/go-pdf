package db

import (
	"context"
	conf "github.com/adnpa/gpdf/config"
	"github.com/go-acme/lego/v4/log"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init 初始化Postgres连接
func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/gpdf?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AllowGlobalUpdate = true
	sqldb, _ := db.DB()
	log.Printf("db:", sqldb)
	log.Println(conf.Cfg.MaxOpenConns)
	log.Println(conf.Cfg.MaxOpenConns)
	sqldb.SetMaxOpenConns(conf.Cfg.MaxOpenConns)
	sqldb.SetMaxIdleConns(conf.Cfg.MaxIdleConns)
	// gorm设置db
	//query.SetDefault(db)
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
