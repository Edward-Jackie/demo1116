package main

import (
	cfg "demo1116/config"
	"demo1116/db"
	"demo1116/logger"
	"demo1116/router"
)

func main() {
	// 初始化日志
	logger.InitLogger("./logs/log")
	// 初始化配置
	cfg.InitConfig()
	// 初始化数据库
	db.InitDb()
	// 初始化web服务
	router.InitWeb()
}
