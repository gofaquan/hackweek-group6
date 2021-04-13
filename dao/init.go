package dao

import (
	"github.com/kashimashino/hackweek-group6/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func InitDB() {
	DB, err = gorm.Open(mysql.Open(os.Getenv("MYSQL")), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	err = DB.AutoMigrate(
		&model.User{},
		&model.Temperalor{},
	)
	if err != nil {
		panic("模型迁移失败")
	}
}
