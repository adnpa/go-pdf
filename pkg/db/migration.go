package db

import (
	"github.com/adnpa/gpdf/app/user/model"
)

func migration() {
	db.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&model.User{})
}
