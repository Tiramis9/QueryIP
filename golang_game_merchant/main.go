package main

import (
	"fmt"
	"golang_game_merchant/global"
	"golang_game_merchant/model"
	"golang_game_merchant/router"
	"net/http"
)

func main() {
	// 初始化配置文件
	global.ConfigInit()

	// todo: 初始化日志

	// 初始化数据库
	model.DbInit()

	// 路由初始化
	engine := router.RouteInit()

	// http服务初始化
	addr := fmt.Sprintf("0.0.0.0:%v", 9876)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
