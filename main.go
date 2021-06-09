package main

import (
	"apihut-server/config"
	"apihut-server/repository/mysql"
	"apihut-server/repository/redis"
	"apihut-server/router"
	"fmt"
)

func main() {
	// 加载配置
	err := config.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 初始化数据库
	if err = mysql.Init(config.Conf.MySQLConfig); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer mysql.Close()
	// 初始化Redis
	if err = redis.Init(); err != nil {
		fmt.Println("redis init err:", err.Error())
		return
	}
	defer redis.Close()

	// 加载路由
	r := router.Setup()

	r.Run()
}
