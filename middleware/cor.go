package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨越中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求来源-所有
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许请求
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许头文件
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
