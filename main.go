package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kashimashino/hackweek-group6/dao"
	"github.com/kashimashino/hackweek-group6/route"
)

func main() {
	dao.InitDB()

	//路由装载
	route.InitRouter()
}
