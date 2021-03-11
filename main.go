package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	//启动数据库
	model.InitDb()

	routes.InitRouter()

}
