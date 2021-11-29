package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func setRouteGroup() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.GET("/login", login)
	v1.GET("submit", submit)

	v2 := r.Group("/v2")
	v2.POST("/login", login)
	v2.POST("/submit", submit)
	return r
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
