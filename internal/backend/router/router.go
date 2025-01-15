package backend

import (
	"demo1116/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.RouterGroup) {
	b := r.Group("/internal")
	// 添加跨域
	b.Use(middleware.Cors())
	{
		b.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Hello User")
		})

		b.DELETE("/user/:id", func(c *gin.Context) {
			p := c.Param("id")
			fmt.Println(p)
		})
	}
}
