package main

import (
	"bookmanager/dao/mysql"
	"bookmanager/router"
)

func main() {
	//初始化数据库
	mysql.InitMysql()
	//初始化路由
	r := router.InitRouter()
	r.Run()
}
