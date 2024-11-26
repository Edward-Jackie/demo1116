package router

import (
	cfg "demo1116/config"
	backend "demo1116/internal/backend/router"
	system "demo1116/internal/system/router"
	log "demo1116/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitWeb() {
	gin.SetMode(gin.ReleaseMode)

	gin.DefaultWriter = log.GinLog
	web := gin.Default()

	// 使用中间件
	web.Use()

	r := web.Group("/api/v1")
	r.Use() // 解决跨域问题
	// 系统路由初始化
	system.InitRouter(r)

	// 业务模块初始化
	backend.InitRouter(r)

	if len(cfg.Cfg.Server.Port) == 0 {
		cfg.Cfg.Server.Port = "8080"
	}
	log.GLog.Info(fmt.Sprintf("项目启动 端口 : %v", cfg.Cfg.Server.Port))
	web.Run(":" + cfg.Cfg.Server.Port)
}
