package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// usable middleware list:
// https://github.com/gin-gonic/contrib/blob/master/README.md

func setGlobalMiddleware() *gin.Engine {
	r := gin.Default()
	// register
	r.Use(Middleware())
	r.GET("/ce", func(c *gin.Context) {
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		c.JSON(200, gin.H{"request": req})
	})
	return r
}

func setLocalMiddleware() *gin.Engine {
	r := gin.Default()
	r.GET("/ce", Middleware(), func(c *gin.Context) {
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		c.JSON(200, gin.H{"request": req})
	})
	return r
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// set key into context
		c.Set("request", "mw")
	}
}
