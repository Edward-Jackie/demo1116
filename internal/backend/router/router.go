package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.RouterGroup) {
	b := r.Group("/")
	{
		b.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Hello User")
		})
	}
}
