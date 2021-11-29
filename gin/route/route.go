package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func setRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func setApiParam() *gin.Engine {
	r := gin.Default()
	r.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		action = strings.Trim(action, "/")
		context.String(http.StatusOK, name+" is "+action)
	})
	return r
}

func setUrlParam() *gin.Engine {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		// name := c.Query("name")
		name := c.DefaultQuery("name", "liu")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	return r
}

func setFormParam() *gin.Engine {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		tp := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.String(http.StatusOK, fmt.Sprintf("username:%s, password:%s, type:%s", username, password, tp))
	})
	return r
}
