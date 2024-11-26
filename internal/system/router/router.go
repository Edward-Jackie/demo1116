package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter 初始化系统服务
func InitRouter(r *gin.RouterGroup) {
	s := r.Group("/sys")
	{
		s.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Hello Admin")
		})
	}
}
