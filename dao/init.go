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
	//dsn := "root:MyNewPass!4@tcp(127.0.0.1:3306)/hackweek?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(os.Getenv("MYSQL")), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	err = DB.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		panic("模型迁移失败")
	}
}
